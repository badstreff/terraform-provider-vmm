package getting

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// VirtualMachineGetter interface for getting clouds
type VirtualMachineGetter interface {
	GetByID(string) (*VirtualMachine, error)
	GetVirtualNetworkAdapters(string) ([]*VirtualNetworkAdapter, error)
	GetAll() ([]*VirtualMachine, error)
}

// VirtualMachineService performs the get operations of the cloud
type VirtualMachineService Service

// NewVirtualMachineService creates a new cloud service
func NewVirtualMachineService(doer doer, serviceURL string, stamp string) *VirtualMachineService {
	url, err := url.Parse(serviceURL)
	if err != nil {
		return nil
	}
	return &VirtualMachineService{doer: doer, serviceURL: *url, stampID: stamp}
}

type virtualMachineFeed struct {
	VirtualMachines []*VirtualMachine `json:"value"`
}

// VirtualMachine object
type VirtualMachine struct {
	AddedTime                      *string                            `json:"AddedTime"`
	Agent                          *bool                              `json:"Agent"`
	AllocatedGPU                   *string                            `json:"AllocatedGPU"`
	BackupEnabled                  *bool                              `json:"BackupEnabled"`
	BlockLiveMigrationIfHostBusy   *bool                              `json:"BlockLiveMigrationIfHostBusy"`
	CanVMConnect                   *bool                              `json:"CanVMConnect"`
	CheckpointLocation             *string                            `json:"CheckpointLocation"`
	CloudID                        *string                            `json:"CloudId" json:"CloudId,omitempty"`
	ComputerName                   *string                            `json:"ComputerName"`
	ComputerTierID                 *string                            `json:"ComputerTierId"`
	CostCenter                     *string                            `json:"CostCenter"`
	CPUCount                       *uint8                             `json:"CPUCount"`
	CPUMax                         *int32                             `json:"CPUMax"`
	CPUReserve                     *int32                             `json:"CPUReserve"`
	CPUType                        *string                            `json:"CPUType"`
	CPUUtilization                 *int32                             `json:"CPUUtilization"`
	CreationSource                 *string                            `json:"CreationSource"`
	CreationTime                   *string                            `json:"CreationTime"`
	DataExchangeEnabled            *bool                              `json:"DataExchangeEnabled"`
	DelayStart                     *int32                             `json:"DelayStart"`
	DeployPath                     *string                            `json:"DeployPath"`
	Description                    *string                            `json:"Description"`
	DiskIO                         *int32                             `json:"DiskIO"`
	Dismiss                        *bool                              `json:"Dismiss"`
	DynamicMemoryDemandMB          *int32                             `json:"DynamicMemoryDemandMB"`
	Enabled                        *bool                              `json:"Enabled"`
	ExcludeFromPRO                 *bool                              `json:"ExcludeFromPRO"`
	ExpectedCPUUtilization         *int32                             `json:"ExpectedCPUUtilization"`
	FailedJobID                    *string                            `json:"FailedJobId"`
	HasPassthroughDisk             *bool                              `json:"HasPassthroughDisk"`
	HasSavedState                  *bool                              `json:"HasSavedState"`
	HasVMAdditions                 *bool                              `json:"HasVMAdditions"`
	HeartbeatEnabled               *bool                              `json:"HeartbeatEnabled"`
	HighlyAvailable                *bool                              `json:"HighlyAvailable"`
	ID                             *string                            `json:"ID"`
	IsFaultTolerant                *bool                              `json:"IsFaultTolerant"`
	IsHighlyAvailable              *bool                              `json:"IsHighlyAvailable"`
	IsUndergoingLiveMigration      *bool                              `json:"IsUndergoingLiveMigration"`
	LastRestoredCheckpointID       *string                            `json:"LastRestoredCheckpointID"`
	LibraryGroup                   *string                            `json:"LibraryGroup"`
	LimitCPUForMigration           *bool                              `json:"LimitCPUForMigration"`
	LimitCPUFunctionality          *bool                              `json:"LimitCPUFunctionality"`
	VMNetworkAssignments           []*VMNetworkAssignment             `json:"VMNetworkAssignments"`
	Location                       *string                            `json:"Location"`
	MarkedAsTemplate               *bool                              `json:"MarkedAsTemplate"`
	Memory                         *int32                             `json:"Memory"`
	DynamicMemoryEnabled           *bool                              `json:"DynamicMemoryEnabled"`
	DynamicMemoryMinimumMB         *int32                             `json:"DynamicMemoryMinimumMB"`
	DynamicMemoryMaximumMB         *int32                             `json:"DynamicMemoryMaximumMB"`
	MemoryAssignedMB               *int32                             `json:"MemoryAssignedMB"`
	MemoryAvailablePercentage      *int16                             `json:"MemoryAvailablePercentage"`
	ModifiedTime                   *string                            `json:"ModifiedTime"`
	MostRecentTaskID               *string                            `json:"MostRecentTaskId"`
	Name                           *string                            `json:"Name"`
	NetworkUtilization             *int32                             `json:"NetworkUtilization"`
	NumLock                        *bool                              `json:"NumLock"`
	OperatingSystem                *string                            `json:"OperatingSystem"`
	OperatingSystemInstance        *OperatingSystem                   `json:"OperatingSystemInstance"`
	OperatingSystemShutdownEnabled *bool                              `json:"OperatingSystemShutdownEnabled"`
	Operation                      *string                            `json:"Operation"`
	Owner                          *UserAndRole                       `json:"Owner"`
	GrantedToList                  []*UserAndRole                     `json:"GrantedToList"`
	Path                           *string                            `json:"Path"`
	PerfCPUUtilization             *int32                             `json:"PerfCPUUtilization"`
	PerfDiskBytesRead              *string                            `json:"PerfDiskBytesRead"`
	PerfDiskBytesWrite             *string                            `json:"PerfDiskBytesWrite"`
	PerfNetworkBytesRead           *string                            `json:"PerfNetworkBytesRead"`
	PerfNetworkBytesWrite          *string                            `json:"PerfNetworkBytesWrite"`
	CPURelativeWeight              *int32                             `json:"CPURelativeWeight"`
	Retry                          *bool                              `json:"Retry"`
	RunGuestAccount                *string                            `json:"RunGuestAccount"`
	ServiceDeploymentErrorMessage  *string                            `json:"ServiceDeploymentErrorMessage"`
	ServiceID                      *string                            `json:"ServiceId"`
	VMShieldingDataID              *string                            `json:"VMShieldingDataId"`
	KeyProtectorOwner              *OwnerOrGuardian                   `json:"KeyProtectorOwner"`
	KeyProtectorGuardians          []*OwnerOrGuardian                 `json:"KeyProtectorGuardians"`
	SecuritySummary                *string                            `json:"SecuritySummary"`
	SharePath                      *string                            `json:"SharePath"`
	Shielded                       *bool                              `json:"Shielded"`
	SourceObjectType               *string                            `json:"SourceObjectType"`
	StartAction                    *string                            `json:"StartAction"`
	StartVM                        *bool                              `json:"StartVM"`
	Status                         *string                            `json:"Status"`
	Statusstring                   *string                            `json:"Statusstring"`
	StopAction                     *string                            `json:"StopAction"`
	Tag                            *string                            `json:"Tag"`
	TimeSynchronizationEnabled     *bool                              `json:"TimeSynchronizationEnabled"`
	TotalSize                      *string                            `json:"TotalSize"`
	Undo                           *bool                              `json:"Undo"`
	UndoDisksEnabled               *bool                              `json:"UndoDisksEnabled"`
	UpgradeDomain                  *int32                             `json:"UpgradeDomain"`
	UseCluster                     *bool                              `json:"UseCluster"`
	UseLAN                         *bool                              `json:"UseLAN"`
	VirtualHardDiskID              *string                            `json:"VirtualHardDiskId"`
	VirtualizationPlatform         *string                            `json:"VirtualizationPlatform"`
	CapabilityProfile              *string                            `json:"CapabilityProfile"`
	VMBaseConfigurationID          *string                            `json:"VMBaseConfigurationId"`
	VMConnection                   *string                            `json:"VMConnection"`
	VMConfigResource               *string                            `json:"VMConfigResource"`
	VMCPath                        *string                            `json:"VMCPath"`
	VMHostName                     *string                            `json:"VMHostName"`
	VMID                           *string                            `json:"VMId"`
	StampID                        *string                            `json:"StampId"`
	VMResource                     *string                            `json:"VMResource"`
	VMResourceGroup                *string                            `json:"VMResourceGroup"`
	VirtualMachineState            *string                            `json:"VirtualMachineState"`
	VMTemplateID                   *string                            `json:"VMTemplateId"`
	HardwareProfileID              *string                            `json:"HardwareProfileId"`
	BlockDynamicOptimization       *bool                              `json:"BlockDynamicOptimization"`
	CPULimitForMigration           *bool                              `json:"CPULimitForMigration"`
	CPULimitFunctionality          *bool                              `json:"CPULimitFunctionality"`
	Domain                         *string                            `json:"Domain"`
	DynamicMemoryBufferPercentage  *int32                             `json:"DynamicMemoryBufferPercentage"`
	FullName                       *string                            `json:"FullName"`
	MemoryWeight                   *int32                             `json:"MemoryWeight"`
	OrganizationName               *string                            `json:"OrganizationName"`
	DelayStartSeconds              *int32                             `json:"DelayStartSeconds"`
	ProductKey                     *string                            `json:"ProductKey"`
	WorkGroup                      *string                            `json:"WorkGroup"`
	TimeZone                       *int32                             `json:"TimeZone"`
	RunAsAccountUserName           *string                            `json:"RunAsAccountUserName"`
	UserName                       *string                            `json:"UserName"`
	Password                       *string                            `json:"Password"`
	LocalAdminRunAsAccountName     *string                            `json:"LocalAdminRunAsAccountName"`
	LocalAdminUserName             *string                            `json:"LocalAdminUserName"`
	LocalAdminPassword             *string                            `json:"LocalAdminPassword"`
	LinuxDomainName                *string                            `json:"LinuxDomainName"`
	LinuxAdministratorSSHKey       *string                            `json:"LinuxAdministratorSSHKey"`
	LinuxAdministratorSSHKeystring *string                            `json:"LinuxAdministratorSSHKeystring"`
	CloudVMRoleName                *string                            `json:"CloudVMRoleName"`
	Generation                     *int32                             `json:"Generation"`
	DeploymentErrorInfo            *ErrorInfo                         `json:"DeploymentErrorInfo"`
	NewVirtualNetworkAdapterInput  []*NewVMVirtualNetworkAdapterInput `json:"NewVirtualNetworkAdapterInput"`
	IsRecoveryVM                   *bool                              `json:"IsRecoveryVM"`
}

// VMNetworkAssignment object
type VMNetworkAssignment struct {
	VirtualNetworkAdapterID string `json:"VirtualNetworkAdapterId"`
	VMNetworkName           string `json:"VMNetworkName"`
}

// NewVMVirtualNetworkAdapterInput object
type NewVMVirtualNetworkAdapterInput struct {
	VMNetworkName   string `json:"VMNetworkName"`
	VMNetworkID     string `json:"VMNetworkId"`
	VMSubnetName    string `json:"VMSubnetName"`
	MACAddressType  string `json:"MACAddressType"`
	MACAddress      string `json:"MACAddress"`
	IPv4AddressType string `json:"IPv4AddressType"`
	IPv6AddressType string `json:"IPv6AddressType"`
	VLanEnabled     bool   `json:"VLanEnabled"`
	VLanID          int16  `json:"VLanId"`
}

// GetByID gets a virtual hard vm by its ID
func (s VirtualMachineService) GetByID(ID string) (*VirtualMachine, error) {
	vm := &VirtualMachine{}
	uri := fmt.Sprintf(s.serviceURL.String()+"(ID=guid'%s',StampId=guid'%s')", ID, s.stampID)
	body, err := get(s, uri)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, vm)
	if err != nil {
		return nil, err
	}
	return vm, nil
}

// GetVirtualNetworkAdapters returns all network adapters for a virtual machine
func (s VirtualMachineService) GetVirtualNetworkAdapters(ID string) ([]*VirtualNetworkAdapter, error) {
	var feed virtualNetworkAdapterFeed
	var adapters []*VirtualNetworkAdapter
	uri := fmt.Sprintf(s.serviceURL.String()+"(ID=guid'%s',StampId=guid'%s')/VirtualNetworkAdapters", ID, s.stampID)
	body, err := get(s, uri)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &feed)
	if err != nil {
		return nil, err
	}
	for _, adapter := range feed.VirtualNetworkAdapters {
		adapters = append(adapters, adapter)
	}
	return adapters, nil
}

// GetAll get all the vm objects
func (s VirtualMachineService) GetAll() ([]*VirtualMachine, error) {
	var feed virtualMachineFeed
	var VirtualMachines []*VirtualMachine
	body, err := get(s, s.serviceURL.String())
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &feed)
	if err != nil {
		return nil, err
	}
	for _, vm := range feed.VirtualMachines {
		VirtualMachines = append(VirtualMachines, vm)
	}
	return VirtualMachines, nil
}
