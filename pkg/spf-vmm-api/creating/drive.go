package creating

import (
	"encoding/json"
	"log"
	"net/url"
)

// VirtualDiskDriveCreator interface for creating virtual disk drives
type VirtualDiskDriveCreator interface {
	Create(*VirtualDiskDrive) (*string, error)
}

// VirtualDiskDriveService performs the create operations
type VirtualDiskDriveService Service

// NewVirtualDiskDriveService creates a new cloud service
func NewVirtualDiskDriveService(doer Doer, serviceURL string, stamp string) *VirtualDiskDriveService {
	url, err := url.Parse(serviceURL)
	if err != nil {
		return nil
	}
	return &VirtualDiskDriveService{Doer: doer, serviceURL: *url, stampID: stamp}
}

// VirtualDiskDrive object for creation
type VirtualDiskDrive struct {
	ODataType         *string `json:"odata.type,omitempty"`
	StampID           *string `json:"StampId,omitempty"`
	VMID              *string `json:"VMId,omitempty"`
	VirtualHardDiskID *string `json:"VirtualHardDiskId,omitempty"`
	Bus               *uint8  `json:",omitempty"`
	LUN               *uint8  `json:",omitempty"`
	IDE               *bool   `json:",omitempty"`
	SCSI              *bool   `json:",omitempty"`
	FileName          *string `json:",omitempty"`
	Path              *string `json:",omitempty"`
}

// Create creates a new virtual disk drive object and returns its ID
func (s VirtualDiskDriveService) Create(drive *VirtualDiskDrive) (id *string, err error) {
	data, err := json.Marshal(*drive)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return create(s.Doer, s.serviceURL.String(), data)
}
