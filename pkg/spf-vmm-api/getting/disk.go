package getting

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// VirtualHardDiskGetter interface for getting clouds
type VirtualHardDiskGetter interface {
	GetByID(string) (*VirtualHardDisk, error)
	GetAll() ([]*VirtualHardDisk, error)
}

// VirtualHardDiskService performs the get operations of the cloud
type VirtualHardDiskService Service

// NewVirtualHardDiskService creates a new cloud service
func NewVirtualHardDiskService(doer doer, serviceURL string, stamp string) *VirtualHardDiskService {
	url, err := url.Parse(serviceURL)
	if err != nil {
		return nil
	}
	return &VirtualHardDiskService{doer: doer, serviceURL: *url, stampID: stamp}
}

type virtualHardDiskFeed struct {
	VirtualHardDisks []*VirtualHardDisk `json:"value"`
}

// VirtualHardDisk object
type VirtualHardDisk struct {
	AddedTime                        *string          `json:"AddedTime"`
	Accessibility                    *string          `json:"Accessibility"`
	Description                      *string          `json:"Description"`
	Directory                        *string          `json:"Directory"`
	Enabled                          *bool            `json:"Enabled"`
	ID                               *string          `json:"ID"`
	JobGroupID                       *string          `json:"JobGroupID"`
	MaximumSize                      *string          `json:"MaximumSize"`
	ModifiedTime                     *string          `json:"ModifiedTime"`
	Name                             *string          `json:"Name"`
	Owner                            *UserAndRole     `json:"Owner"`
	OwnerSid                         *string          `json:"OwnerSid"`
	ParentDiskID                     *string          `json:"ParentDiskID"`
	SharePath                        *string          `json:"SharePath"`
	Shielded                         *bool            `json:"Shielded"`
	Size                             *string          `json:"Size"`
	State                            *string          `json:"State"`
	VHDType                          *string          `json:"VHDType"`
	VMID                             *string          `json:"VMID"`
	TemplateID                       *string          `json:"TemplateID"`
	StampID                          *string          `json:"StampID"`
	FamilyName                       *string          `json:"FamilyName"`
	Release                          *string          `json:"Release"`
	CloudID                          *string          `json:"CloudID"`
	HostVolumeID                     *string          `json:"HostVolumeID"`
	IsOrphaned                       *bool            `json:"IsOrphaned"`
	IsResourceGroup                  *bool            `json:"IsResourceGroup"`
	LibraryGroup                     *string          `json:"LibraryGroup"`
	LibraryShareID                   *string          `json:"LibraryShareID"`
	Location                         *string          `json:"Location"`
	Namespace                        *string          `json:"Namespace"`
	ReleaseTime                      *string          `json:"ReleaseTime"`
	SANCopyCapable                   *bool            `json:"SANCopyCapable"`
	Type                             *string          `json:"Type"`
	VirtualizationPlatform           *string          `json:"VirtualizationPlatform"`
	OperatingSystem                  *string          `json:"OperatingSystem"`
	OperatingSystemInstance          *OperatingSystem `json:"OperatingSystemInstance"`
	OperatingSystemID                *string          `json:"OperatingSystemID"`
	Tag                              []*string        `json:"Tag"`
	VolumeSignatureCatalogName       *string          `json:"VolumeSignatureCatalogName"`
	VolumeSignatureCatalogVersion    *string          `json:"VolumeSignatureCatalogVersion"`
	VolumeSignatureCatalogIssuerName *string          `json:"VolumeSignatureCatalogIssuerName"`
}

// GetByID gets a virtual hard disk by its ID
func (s VirtualHardDiskService) GetByID(ID string) (*VirtualHardDisk, error) {
	disk := &VirtualHardDisk{}
	uri := fmt.Sprintf(s.serviceURL.String()+"(ID=guid'%s',StampId=guid'%s')", ID, s.stampID)
	body, err := get(s, uri)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, disk)
	if err != nil {
		return nil, err
	}
	return disk, nil
}

// GetAll get all the virtual hard disk objects
func (s VirtualHardDiskService) GetAll() ([]*VirtualHardDisk, error) {
	var feed virtualHardDiskFeed
	var VirtualHardDisks []*VirtualHardDisk
	body, err := get(s, s.serviceURL.String())
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &feed)
	if err != nil {
		return nil, err
	}
	for _, disk := range feed.VirtualHardDisks {
		VirtualHardDisks = append(VirtualHardDisks, disk)
	}
	return VirtualHardDisks, nil
}
