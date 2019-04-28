package getting

import (
	"reflect"
	"testing"
)

func TestVirtualMachineGetByID(t *testing.T) {
	doer := MockDoer{}
	doer.StatusCode = 200
	doer.Data = `{
    "AddedTime": "2019-04-27T20:23:09.783-05:00",
    "Agent": null,
    "AllocatedGPU": null,
    "BackupEnabled": true,
    "BlockDynamicOptimization": null,
    "BlockLiveMigrationIfHostBusy": null,
    "CPUCount": 8,
    "CPULimitForMigration": null,
    "CPULimitFunctionality": null,
    "CPUMax": 100,
    "CPURelativeWeight": null,
    "CPUReserve": 0,
    "CPUType": "3.60 GHz Xeon (2 MB L2 cache)",
    "CPUUtilization": 0,
    "CanVMConnect": false,
    "CapabilityProfile": "Hyper-V",
    "CheckpointLocation": null,
    "CloudId": "26782fc4-66cc-4233-9b29-b0533723b657",
    "CloudVMRoleName": null,
    "ComputerName": null,
    "ComputerTierId": null,
    "CostCenter": null,
    "CreationSource": "Temporary Template36a671c0-5218-471f-ae0e-c32b34a477be",
    "CreationTime": "2019-04-27T20:23:09.783-05:00",
    "DataExchangeEnabled": true,
    "DelayStart": 0,
    "DelayStartSeconds": null,
    "DeployPath": null,
    "DeploymentErrorInfo": {
        "CloudProblem": null,
        "Code": null,
        "DetailedCode": null,
        "DetailedErrorCode": null,
        "DetailedSource": null,
        "DisplayableErrorCode": null,
        "ErrorCodeString": null,
        "ErrorType": null,
        "ExceptionDetails": null,
        "IsConditionallyTerminating": null,
        "IsDeploymentBlocker": null,
        "IsMomAlert": null,
        "IsSuccess": null,
        "IsTerminating": null,
        "MessageParameters": null,
        "MomAlertSeverity": null,
        "Problem": null,
        "RecommendedAction": null,
        "RecommendedActionCLI": null,
        "ShowDetailedError": null
    },
    "Description": "",
    "DiskIO": 0,
    "Dismiss": null,
    "Domain": null,
    "DynamicMemoryBufferPercentage": null,
    "DynamicMemoryDemandMB": 0,
    "DynamicMemoryEnabled": false,
    "DynamicMemoryMaximumMB": null,
    "DynamicMemoryMinimumMB": null,
    "Enabled": true,
    "ExcludeFromPRO": false,
    "ExpectedCPUUtilization": 20,
    "FailedJobID": null,
    "FullName": null,
    "Generation": 2,
    "GrantedToList": [],
    "HardwareProfileId": null,
    "HasPassthroughDisk": false,
    "HasSavedState": false,
    "HasVMAdditions": false,
    "HeartbeatEnabled": true,
    "HighlyAvailable": null,
    "ID": "36397a12-3b2e-42b4-afdc-ddf4b55fd743",
    "IsFaultTolerant": false,
    "IsHighlyAvailable": true,
    "IsRecoveryVM": false,
    "IsUndergoingLiveMigration": false,
    "KeyProtectorGuardians": [],
    "KeyProtectorOwner": {
        "EncryptionCertificate": null,
        "Name": null,
        "SigningCertificate": null
    },
    "LastRestoredCheckpointId": "80c77f11-2ac1-417c-bcde-d1feed33a9fb",
    "LibraryGroup": "",
    "LimitCPUForMigration": false,
    "LimitCPUFunctionality": false,
    "LinuxAdministratorSSHKey": null,
    "LinuxAdministratorSSHKeyString": null,
    "LinuxDomainName": null,
    "LocalAdminPassword": null,
    "LocalAdminRunAsAccountName": null,
    "LocalAdminUserName": null,
    "Location": "",
    "MarkedAsTemplate": false,
    "Memory": 4096,
    "MemoryAssignedMB": 0,
    "MemoryAvailablePercentage": null,
    "MemoryWeight": 5000,
    "ModifiedTime": "2019-04-29T23:06:01.2823063-05:00",
    "MostRecentTaskId": "64e111c8-6a41-4ec0-a044-00b87e9fb8c2",
    "Name": "TerraformSPFTestVM",
    "NetworkUtilization": 0,
    "NewVirtualNetworkAdapterInput": [],
    "NumLock": null,
    "OperatingSystem": "Unknown",
    "OperatingSystemInstance": {
        "Architecture": "x86",
        "Description": "Virtual Machine Manager was unable to determine the operating system for this object.",
        "Edition": null,
        "Name": "Unknown",
        "OSType": "Other",
        "ProductType": null,
        "Version": null
    },
    "OperatingSystemShutdownEnabled": true,
    "Operation": null,
    "OrganizationName": null,
    "Owner": {
        "RoleID": "18f3f16d-658a-4acc-ad55-4b1a3b55fc00",
        "RoleName": "Administrator",
        "UserName": "EXAMPLE.COM\\ExampleUser"
    },
    "Password": null,
    "Path": null,
    "PerfCPUUtilization": -1,
    "PerfDiskBytesRead": "-1",
    "PerfDiskBytesWrite": "-1",
    "PerfNetworkBytesRead": "-1",
    "PerfNetworkBytesWrite": "-1",
    "ProductKey": null,
    "Retry": null,
    "RunAsAccountUserName": null,
    "RunGuestAccount": null,
    "SecuritySummary": "None",
    "ServiceDeploymentErrorMessage": null,
    "ServiceId": null,
    "SharePath": null,
    "Shielded": null,
    "SourceObjectType": "VM Template",
    "StampId": "2b085433-709f-4ebe-81fb-c1af3ffa9fd1",
    "StartAction": "TurnOnVMIfRunningWhenVSStopped",
    "StartVM": null,
    "Status": "PowerOff",
    "StatusString": "Stopped",
    "StopAction": "SaveVM",
    "Tag": "(none)",
    "TimeSynchronizationEnabled": true,
    "TimeZone": null,
    "TotalSize": "4194304",
    "Undo": null,
    "UndoDisksEnabled": false,
    "UpgradeDomain": null,
    "UseCluster": null,
    "UseLAN": null,
    "UserName": null,
    "VMBaseConfigurationId": null,
    "VMCPath": null,
    "VMConfigResource": null,
    "VMConnection@odata.mediaContentType": "application/x-rdp",
    "VMHostName": null,
    "VMId": "80c77f11-2ac1-417c-bcde-d1feed33a9fb",
    "VMNetworkAssignments": [],
    "VMResource": null,
    "VMResourceGroup": null,
    "VMShieldingDataId": null,
    "VMTemplateId": null,
    "VirtualHardDiskId": null,
    "VirtualMachineState": "PowerOff",
    "VirtualizationPlatform": "HyperV",
    "WorkGroup": null,
    "odata.metadata": "https://VMM.EXAMPLE.COM:8090/SC2012R2/VMM/Microsoft.Management.Odata.svc/$metadata#VirtualMachines/@Element"
}`
	vmService := NewVirtualMachineService(doer, "", "")
	vm, err := vmService.GetByID("")
	if err != nil {
		t.Errorf("Error attempting to parse vm")
		t.Errorf("%v", err)
	}
	expected := &VirtualMachine{
		AddedTime:                     sPtr("2019-04-27T20:23:09.783-05:00"),
		BackupEnabled:                 bPtr(true),
		CPUCount:                      uint8Ptr(8),
		CPUMax:                        int32Ptr(100),
		CPUReserve:                    int32Ptr(0),
		CPUType:                       sPtr("3.60 GHz Xeon (2 MB L2 cache)"),
		CPUUtilization:                int32Ptr(0),
		CanVMConnect:                  bPtr(false),
		CapabilityProfile:             sPtr("Hyper-V"),
		CloudID:                       sPtr("26782fc4-66cc-4233-9b29-b0533723b657"),
		CreationSource:                sPtr("Temporary Template36a671c0-5218-471f-ae0e-c32b34a477be"),
		CreationTime:                  sPtr("2019-04-27T20:23:09.783-05:00"),
		DataExchangeEnabled:           bPtr(true),
		DelayStart:                    int32Ptr(0),
		DeploymentErrorInfo:           &ErrorInfo{},
		Description:                   sPtr(""),
		DiskIO:                        int32Ptr(0),
		DynamicMemoryDemandMB:         int32Ptr(0),
		DynamicMemoryEnabled:          bPtr(false),
		Enabled:                       bPtr(true),
		ExcludeFromPRO:                bPtr(false),
		ExpectedCPUUtilization:        int32Ptr(20),
		Generation:                    int32Ptr(2),
		GrantedToList:                 make([]*UserAndRole, 0),
		HasPassthroughDisk:            bPtr(false),
		HasSavedState:                 bPtr(false),
		HasVMAdditions:                bPtr(false),
		HeartbeatEnabled:              bPtr(true),
		ID:                            sPtr("36397a12-3b2e-42b4-afdc-ddf4b55fd743"),
		IsFaultTolerant:               bPtr(false),
		IsHighlyAvailable:             bPtr(true),
		IsRecoveryVM:                  bPtr(false),
		IsUndergoingLiveMigration:     bPtr(false),
		KeyProtectorGuardians:         make([]*OwnerOrGuardian, 0),
		KeyProtectorOwner:             &OwnerOrGuardian{},
		LastRestoredCheckpointID:      sPtr("80c77f11-2ac1-417c-bcde-d1feed33a9fb"),
		LibraryGroup:                  sPtr(""),
		LimitCPUForMigration:          bPtr(false),
		LimitCPUFunctionality:         bPtr(false),
		Location:                      sPtr(""),
		MarkedAsTemplate:              bPtr(false),
		Memory:                        int32Ptr(4096),
		MemoryAssignedMB:              int32Ptr(0),
		MemoryWeight:                  int32Ptr(5000),
		ModifiedTime:                  sPtr("2019-04-29T23:06:01.2823063-05:00"),
		MostRecentTaskID:              sPtr("64e111c8-6a41-4ec0-a044-00b87e9fb8c2"),
		Name:                          sPtr("TerraformSPFTestVM"),
		NetworkUtilization:            int32Ptr(0),
		NewVirtualNetworkAdapterInput: make([]*NewVMVirtualNetworkAdapterInput, 0),
		OperatingSystem:               sPtr("Unknown"),
		OperatingSystemInstance: &OperatingSystem{
			Architecture: sPtr("x86"),
			Description:  sPtr("Virtual Machine Manager was unable to determine the operating system for this object."),
			Name:         sPtr("Unknown"),
			OSType:       sPtr("Other"),
		},
		OperatingSystemShutdownEnabled: bPtr(true),
		Owner: &UserAndRole{
			RoleID:   "18f3f16d-658a-4acc-ad55-4b1a3b55fc00",
			RoleName: "Administrator",
			UserName: "EXAMPLE.COM\\ExampleUser",
		},
		PerfCPUUtilization:         int32Ptr(-1),
		PerfDiskBytesRead:          sPtr("-1"),
		PerfDiskBytesWrite:         sPtr("-1"),
		PerfNetworkBytesRead:       sPtr("-1"),
		PerfNetworkBytesWrite:      sPtr("-1"),
		SecuritySummary:            sPtr("None"),
		SourceObjectType:           sPtr("VM Template"),
		StampID:                    sPtr("2b085433-709f-4ebe-81fb-c1af3ffa9fd1"),
		StartAction:                sPtr("TurnOnVMIfRunningWhenVSStopped"),
		Status:                     sPtr("PowerOff"),
		Statusstring:               sPtr("Stopped"),
		StopAction:                 sPtr("SaveVM"),
		Tag:                        sPtr("(none)"),
		TimeSynchronizationEnabled: bPtr(true),
		TotalSize:                  sPtr("4194304"),
		UndoDisksEnabled:           bPtr(false),
		VMID:                       sPtr("80c77f11-2ac1-417c-bcde-d1feed33a9fb"),
		VMNetworkAssignments:       make([]*VMNetworkAssignment, 0),
		VirtualMachineState:        sPtr("PowerOff"),
		VirtualizationPlatform:     sPtr("HyperV"),
	}
	if !reflect.DeepEqual(vm, expected) {
		t.Errorf("vms not equal")
	}
}

//func TestVirtualMachineGetAll(t *testing.T) {
//	doer := MockDoer{}
//	doer.Data = ``
//	vmService := NewVirtualMachineService(doer, "", "")
//	VMs, err := vmService.GetAll()
//	if err != nil {
//		t.Errorf("Error attempting to parse vm feed")
//		t.Errorf("%v", err)
//	}
//	if len(VMs) != 2 {
//		t.Errorf("Wrong number of vms returned: expected 2, got %v", len(VMs))
//	}
//}
