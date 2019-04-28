package client

import (
	"errors"
	"github.com/badstreff/terraform-provider-vmm/pkg/spf-vmm-api/getting"
)

// CloudService service
type CloudService struct {
	getting.CloudGetter
}

// GetByName returns a cloud object given its name, returns an error if the name matches more that one or not found
func (s *CloudService) GetByName(name string) (*getting.Cloud, error) {
	count := 0
	index := 0
	items, err := s.GetAll()
	if err != nil {
		return nil, err
	}
	for i, item := range items {
		if name == *item.Name {
			index = i
			count++
		}
	}
	if count == 0 {
		return nil, errors.New("Cloud not found")
	} else if count > 1 {
		return nil, errors.New("More than one Cloud found by this name, use GetByID instead")
	}
	return items[index], nil
}
