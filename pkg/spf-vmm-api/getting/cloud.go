package getting

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// CloudGetter interface for getting clouds
type CloudGetter interface {
	GetByID(string) (*Cloud, error)
	GetAll() ([]*Cloud, error)
}

// CloudService performs the get operations of the cloud
type CloudService Service

// NewCloudService creates a new cloud service
func NewCloudService(doer doer, serviceURL string, stamp string) *CloudService {
	url, err := url.Parse(serviceURL)
	if err != nil {
		return nil
	}
	return &CloudService{doer: doer, serviceURL: *url, stampID: stamp}
}

type cloudFeed struct {
	Clouds []*Cloud `json:"value"`
}

// Cloud cloud object
type Cloud struct {
	ID                      *string `json:"ID"`
	Name                    *string `json:"Name"`
	Description             *string `json:"Description"`
	LastModifiedDate        *string `json:"LastModifiedDate"`
	WritableLibraryPath     *string `json:"WritableLibraryPath"`
	UserRoleID              *string `json:"UserRoleID"`
	StampID                 *string `json:"StampId"`
	ShieldedVMSupportPolicy *string `json:"ShieldedVMSupportPolicy"`
}

// GetByID gets a cloud by its ID
func (s CloudService) GetByID(ID string) (*Cloud, error) {
	cloud := &Cloud{}
	uri := fmt.Sprintf(s.serviceURL.String()+"(ID=guid'%s',StampId=guid'%s')", ID, s.stampID)
	body, err := get(s, uri)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, cloud)
	if err != nil {
		return nil, err
	}
	return cloud, nil
}

// GetAll get all the cloud objects
func (s CloudService) GetAll() ([]*Cloud, error) {
	var Clouds []*Cloud
	var feed cloudFeed
	body, err := get(s, s.serviceURL.String())
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &feed)
	if err != nil {
		return nil, err
	}
	for _, cloud := range feed.Clouds {
		Clouds = append(Clouds, cloud)
	}
	return Clouds, nil
}
