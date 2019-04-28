package client

import (
	"errors"
	"github.com/badstreff/terraform-provider-vmm/pkg/spf-vmm-api/getting"
)

// VirtualHardDiskService service
type VirtualHardDiskService struct {
	getting.VirtualHardDiskGetter
}

// GetByName returns a virtual hard disk object given its name, returns an error if the name matches more that one or not found
func (s *VirtualHardDiskService) GetByName(name string) (*getting.VirtualHardDisk, error) {
	count := 0
	index := 0
	disks, err := s.GetAll()
	if err != nil {
		return nil, err
	}
	for i, disk := range disks {
		if name == *disk.Name {
			index = i
			count++
		}
	}
	if count == 0 {
		return nil, errors.New("VirtalHardDisk not found")
	} else if count > 1 {
		return nil, errors.New("More than on VirtualHardDIsk found by this name, use GetByID instead")
	}
	return disks[index], nil
}
