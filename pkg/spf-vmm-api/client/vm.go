package client

import (
	"errors"
	"github.com/badstreff/terraform-provider-vmm/pkg/spf-vmm-api/creating"
	"github.com/badstreff/terraform-provider-vmm/pkg/spf-vmm-api/deleting"
	"github.com/badstreff/terraform-provider-vmm/pkg/spf-vmm-api/getting"
	"github.com/badstreff/terraform-provider-vmm/pkg/spf-vmm-api/updating"
)

// VirtualMachineService service
type VirtualMachineService struct {
	creating.VirtualMachineCreator
	deleting.GenericDeleter
	getting.VirtualMachineGetter
	updating.VirtualMachineUpdater
}

// Start a virtual machine by its ID
func (s *VirtualMachineService) Start(id string) error {
	return s.performOperation(id, "Start")
}

// Stop a virtual machine by its ID
func (s *VirtualMachineService) Stop(id string) error {
	return s.performOperation(id, "Stop")
}

// Repair a virtual machine by its ID
func (s *VirtualMachineService) Repair(id string) error {
	return s.performOperation(id, "Repair")
}

func (s *VirtualMachineService) performOperation(id string, operation string) error {
	vm, err := s.GetByID(id)
	if err != nil {
		return err
	}
	sPtr := func(s string) *string { return &s }
	data := updating.VirtualMachine{
		ODataType: sPtr("VMM.VirtualMachine"),
		Operation: &operation,
		ID:        &id,
		StampID:   vm.StampID,
	}
	return s.Update(&data)
}

// GetByName returns a virtual machine object given its name, returns an error if the name matches more that one VM or not found
func (s *VirtualMachineService) GetByName(name string) (*getting.VirtualMachine, error) {
	count := 0
	index := 0
	vms, err := s.GetAll()
	if err != nil {
		return nil, err
	}
	for i, vm := range vms {
		if name == *vm.Name {
			index = i
			count++
		}
	}
	if count == 0 {
		return nil, errors.New("VM not found")
	} else if count > 1 {
		return nil, errors.New("More than on VM found by this name, use GetByID instead")
	}
	return vms[index], nil
}
