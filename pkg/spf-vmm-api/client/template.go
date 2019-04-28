package client

import (
	"errors"
	"github.com/badstreff/terraform-provider-vmm/pkg/spf-vmm-api/getting"
)

// VMTemplateService service
type VMTemplateService struct {
	getting.VMTemplateGetter
}

// GetByName returns a vm template object given its name, returns an error if the name matches more that one or not found
func (s *VMTemplateService) GetByName(name string) (*getting.VMTemplate, error) {
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
		return nil, errors.New("Template not found")
	} else if count > 1 {
		return nil, errors.New("More than one Template found by this name, use GetByID instead")
	}
	return items[index], nil
}
