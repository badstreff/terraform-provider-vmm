package creating

import (
	"encoding/json"
	"log"
	"net/url"
)

// VirtualMachineCreator interface for creating virtual machines
type VirtualMachineCreator interface {
	Create(*VirtualMachine) (*string, error)
}

// VirtualMachineService performs the create operations
type VirtualMachineService Service

// NewVirtualMachineService creates a new cloud service
func NewVirtualMachineService(doer Doer, serviceURL string, stamp string) *VirtualMachineService {
	url, err := url.Parse(serviceURL)
	if err != nil {
		return nil
	}
	return &VirtualMachineService{Doer: doer, serviceURL: *url, stampID: stamp}
}

// VirtualMachine object for creation
type VirtualMachine struct {
	ODataType                      *string `json:"odata.type"`
	StampID                        *string `json:"StampId"`
	Name                           *string
	CloudID                        *string `json:"CloudId"`
	VMTemplateID                   *string `json:"VMTemplateId"`
	VirtualHardDiskID              *string `json:"VirtualHardDiskId"`
	HardwareProfileID              *string `json:"HardwareProfileId"`
	Description                    *string
	CostCenter                     *string
	Tag                            *string
	ComputerName                   *string
	BlockDynamicOptimization       *bool
	CPULimitForMigration           *bool
	CPULimitFunctionality          *bool
	CPURelativeWeight              *int
	DelayStartSeconds              *int
	Domain                         *string
	DynamicMemoryBufferPercentage  *int
	DynamicMemoryEnabled           *bool
	DynamicMemoryMaximumMB         *int
	FullName                       *string
	Memory                         *int
	MemoryWeight                   *int
	OrganizationName               *string
	StartAction                    *string
	StartVM                        *bool
	StopAction                     *string
	CPUCount                       *uint8
	Owner                          *UserAndRole
	ProductKey                     *string
	WorkGroup                      *string
	TimeZone                       *int
	RunAsAccountUserName           *string
	UserName                       *string
	Password                       *string
	LocalAdminRunAsAccountName     *string
	LocalAdminUserName             *string
	LocalAdminPassword             *string
	NewVirtualNetworkAdapterInput  []*NewVMVirtualNetworkAdapterInput
	LinuxAdministratorSSHKey       *string
	LinuxDomainName                *string
	LinuxAdministratorSSHKeyString *string
}

// UserAndRole object
type UserAndRole struct {
	UserName *string
	RoleName *string
	RoleID   *string
}

// NewVMVirtualNetworkAdapterInput object
type NewVMVirtualNetworkAdapterInput struct {
	VMNetworkName   *string
	VMNetworkID     *string `json:"VMNetworkId"`
	VMSubnetName    *string
	MACAddressType  *string
	MACAddress      *string
	IPv4AddressType *string
	IPv6AddressType *string
	VLanEnabled     *bool
	VLanID          *int16 `json:"VLanId"`
}

// Create creates a new virtual machine object and returns its ID
func (s VirtualMachineService) Create(vm *VirtualMachine) (id *string, err error) {
	data, err := json.Marshal(*vm)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return create(s.Doer, s.serviceURL.String(), data)
}
