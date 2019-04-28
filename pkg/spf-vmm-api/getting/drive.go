package getting

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// VirtualDiskDriveGetter interface for getting drives
type VirtualDiskDriveGetter interface {
	GetByID(string) (*VirtualDiskDrive, error)
	GetAll() ([]*VirtualDiskDrive, error)
}

// VirtualDiskDriveService performs the get operations of the drive
type VirtualDiskDriveService Service

// NewVirtualDiskDriveService creates a new drive service
func NewVirtualDiskDriveService(doer doer, serviceURL string, stamp string) *VirtualDiskDriveService {
	url, err := url.Parse(serviceURL)
	if err != nil {
		return nil
	}
	return &VirtualDiskDriveService{doer: doer, serviceURL: *url, stampID: stamp}
}

type virtualDiskDriveFeed struct {
	VirtualDiskDrives []*VirtualDiskDrive `json:"value"`
}

// VirtualDiskDrive object
type VirtualDiskDrive struct {
	StampID           *string `json:"StampId"`
	ID                *string `json:"ID"`
	Bus               *uint8  `json:"Bus"`
	BusType           *string `json:"BusType"`
	IsVHD             *bool   `json:"IsVHD"`
	LUN               *uint8  `json:"LUN"`
	Name              *string `json:"Name"`
	VMID              *string `json:"VMId"`
	TemplateID        *string `json:"TemplateId"`
	IsoID             *string `json:"ISOId"`
	HostDrive         *string `json:"HostDrive"`
	IsoLinked         *bool   `json:"ISOLinked"`
	Accessibility     *string `json:"Accessibility"`
	Description       *string `json:"Description"`
	AddedTime         *string `json:"AddedTime"`
	ModifiedTime      *string `json:"ModifiedTime"`
	Enabled           *bool   `json:"Enabled"`
	VirtualHardDiskID *string `json:"VirtualHardDiskId"`
	VolumeType        *string `json:"VolumeType"`
	IDE               *bool   `json:"IDE"`
	SCSI              *bool   `json:"SCSI"`
	FileName          *string `json:"FileName"`
	Path              *string `json:"Path"`
	Size              *string `json:"Size"`
}

// GetByID gets a virtual disk drive by its ID
func (s VirtualDiskDriveService) GetByID(ID string) (*VirtualDiskDrive, error) {
	disk := &VirtualDiskDrive{}
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

// GetAll get all the virtual disk drive objects
func (s VirtualDiskDriveService) GetAll() ([]*VirtualDiskDrive, error) {
	var VirtualDiskDrives []*VirtualDiskDrive
	var feed virtualDiskDriveFeed
	body, err := get(s, s.serviceURL.String())
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &feed)
	if err != nil {
		return nil, err
	}
	for _, disk := range feed.VirtualDiskDrives {
		VirtualDiskDrives = append(VirtualDiskDrives, disk)
	}
	return VirtualDiskDrives, nil
}
