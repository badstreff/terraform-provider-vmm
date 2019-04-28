package updating

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

// JobUpdater interface for updating virtual machines
type JobUpdater interface {
	Update(*Job) error
}

// JobService performs the create operations
type JobService Service

// NewJobService creates a new cloud service
func NewJobService(doer doer, serviceURL string, stamp string) *JobService {
	url, err := url.Parse(serviceURL)
	if err != nil {
		return nil
	}
	return &JobService{doer: doer, serviceURL: *url, stampID: stamp}
}

// Job object for creation
type Job struct {
	ODataType          *string `json:"odata.type"`
	ID                 *string `json:",omitempty"`
	StampID            *string `json:",omitempty"`
	UserName           *string `json:",omitempty"`
	Password           *string `json:",omitempty"`
	SkipLastFailedStep *bool   `json:",omitempty"`
}

// Update updates a virtual machine object
func (s JobService) Update(job *Job) error {
	uri := fmt.Sprintf(s.serviceURL.String()+"(ID=guid'%s',StampId=guid'%s')", *job.ID, s.stampID)
	data, err := json.Marshal(*job)
	if err != nil {
		log.Print(err)
		return err
	}
	return update(s.doer, uri, data)
}
