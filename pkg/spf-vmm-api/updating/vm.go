package updating

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

// VirtualMachineUpdater interface for updating virtual machines
type VirtualMachineUpdater interface {
	Update(*VirtualMachine) error
}

// VirtualMachineService performs the create operations
type VirtualMachineService Service

// NewVirtualMachineService creates a new cloud service
func NewVirtualMachineService(doer doer, serviceURL string, stamp string) *VirtualMachineService {
	url, err := url.Parse(serviceURL)
	if err != nil {
		return nil
	}
	return &VirtualMachineService{doer: doer, serviceURL: *url, stampID: stamp}
}

// VirtualMachine object for creation
type VirtualMachine struct {
	ODataType            *string                 `json:"odata.type"`
	StampID              *string                 `json:"StampId"`
	ID                   *string                 `json:",omitempty"`
	CostCenter           *string                 `json:",omitempty"`
	CPUCount             *uint8                  `json:",omitempty"`
	CPUMax               *int                    `json:",omitempty"`
	CPUReserve           *int                    `json:",omitempty"`
	DelayStartSeconds    *int                    `json:",omitempty"`
	DeployPath           *string                 `json:",omitempty"`
	DiskIO               *int                    `json:",omitempty"`
	Dismiss              *bool                   `json:",omitempty"`
	VMNetworkAssignments []*VMNetworkAssignments `json:",omitempty"`
	Name                 *string                 `json:",omitempty"`
	NetworkUtilization   *int                    `json:",omitempty"`
	Owner                *UserAndRole            `json:",omitempty"`
	GrantedToList        *string                 `json:",omitempty"`
	Tag                  *string                 `json:",omitempty"`
	Operation            *string                 `json:",omitempty"`
	HighlyAvailable      *bool                   `json:",omitempty"`
	Memory               *int                    `json:",omitempty"`
	NumLock              *bool                   `json:",omitempty"`
	Retry                *bool                   `json:",omitempty"`
	Description          *string                 `json:",omitempty"`
	SharePath            *string                 `json:",omitempty"`
	StartAction          *string                 `json:",omitempty"`
	StartVM              *bool                   `json:",omitempty"`
	StopAction           *string                 `json:",omitempty"`
	Undo                 *bool                   `json:",omitempty"`
	VMHostName           *string                 `json:",omitempty"`
	CapabilityProfile    *string                 `json:",omitempty"`
	Agent                *bool                   `json:",omitempty"`
	RelativeWeight       *int                    `json:",omitempty"`
}

// UserAndRole object
type UserAndRole struct {
	UserName *string
	RoleName *string
	RoleID   *string
}

//VMNetworkAssignments object
type VMNetworkAssignments struct {
}

// Update updates a virtual machine object
func (s VirtualMachineService) Update(vm *VirtualMachine) error {
	uri := fmt.Sprintf(s.serviceURL.String()+"(ID=guid'%s',StampId=guid'%s')", *vm.ID, s.stampID)
	data, err := json.Marshal(*vm)
	if err != nil {
		log.Print(err)
		return err
	}
	return update(s.doer, uri, data)
}
