package getting

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// JobGetter interface for getting jobs
type JobGetter interface {
	GetByID(string) (*Job, error)
	GetAll() ([]*Job, error)
}

// JobService performs the get operations of the jobs
type JobService Service

// NewJobService creates a new job service
func NewJobService(doer doer, serviceURL string, stamp string) *JobService {
	url, err := url.Parse(serviceURL)
	if err != nil {
		return nil
	}
	return &JobService{doer: doer, serviceURL: *url, stampID: stamp}
}

type jobFeed struct {
	Jobs []*Job `json:"value"`
}

// Job represents a job object
type Job struct {
	ID                       *string         `json:"ID"`
	Name                     *string         `json:"Name"`
	CurrentStepID            *string         `json:"CurrentStepId"`
	RootStepID               *string         `json:"RootStepId"`
	CmdletName               *string         `json:"CmdletName"`
	EndTime                  *string         `json:"EndTime"`
	IsCompleted              *bool           `json:"IsCompleted"`
	IsRestartable            *bool           `json:"IsRestartable"`
	IsStoppable              *bool           `json:"IsStoppable"`
	IsVisible                *bool           `json:"IsVisible"`
	Owner                    *string         `json:"Owner"`
	OwnerSID                 *string         `json:"OwnerSId"`
	ProgressValue            *int32          `json:"ProgressValue"`
	Progress                 *string         `json:"Progress"`
	ResultObjectTypeName     *string         `json:"ResultObjectTypeName"`
	ResultName               *string         `json:"ResultName"`
	Source                   *string         `json:"Source"`
	StartTime                *string         `json:"StartTime"`
	Status                   *string         `json:"Status"`
	Statusstring             *string         `json:"Statusstring"`
	Target                   *string         `json:"Target"`
	TargetObjectID           *string         `json:"TargetObjectId"`
	Description              *JobDescription `json:"Description"`
	WasNotifiedOfCancel      *bool           `json:"WasNotifiedOfCancel"`
	PROTipID                 *string         `json:"PROTipId"`
	ResultObjectID           *string         `json:"ResultObjectId"`
	TargetObjectType         *string         `json:"TargetObjectType"`
	AreAuditRecordsAvailable *bool           `json:"AreAuditRecordsAvailable"`
	ResultObjectType         *string         `json:"ResultObjectType"`
	StampID                  *string         `json:"StampId"`
	StepsLoaded              *bool           `json:"StepsLoaded"`
	UserName                 *string         `json:"UserName"`
	Password                 *string         `json:"Password"`
	SkipLastFailedStep       *bool           `json:"SkipLastFailedStep"`
	ErrorInfo                *ErrorInfo      `json:"ErrorInfo"`
	AdditionalMessages       []*ErrorInfo    `json:"AdditionalMessages"`
}

// JobDescription metadata about a job
type JobDescription struct {
	CanSkipLastFailedStep         bool   `json:"CanSkipLastFailedStep"`
	Code                          string `json:"Code"`
	DescriptionCodeString         string `json:"DescriptionCodeString"`
	IsRestartable                 bool   `json:"IsRestartable"`
	IsStoppable                   bool   `json:"IsStoppable"`
	Name                          string `json:"Name"`
	RequiresCredentialsForRestart bool   `json:"RequiresCredentialsForRestart"`
}

// GetByID gets a job by its ID
func (s JobService) GetByID(ID string) (*Job, error) {
	job := &Job{}
	uri := fmt.Sprintf(s.serviceURL.String()+"(ID=guid'%s',StampId=guid'%s')", ID, s.stampID)
	body, err := get(s, uri)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, job)
	if err != nil {
		return nil, err
	}
	return job, nil
}

// GetAll get all the job objects
func (s JobService) GetAll() ([]*Job, error) {
	var feed jobFeed
	var Jobs []*Job
	body, err := get(s, s.serviceURL.String())
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &feed)
	if err != nil {
		return nil, err
	}
	for _, job := range feed.Jobs {
		Jobs = append(Jobs, job)
	}
	return Jobs, nil
}
