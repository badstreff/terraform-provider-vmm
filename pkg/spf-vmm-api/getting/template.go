package getting

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// VMTemplateGetter interface for getting vm templates
type VMTemplateGetter interface {
	GetByID(string) (*VMTemplate, error)
	GetAll() ([]*VMTemplate, error)
}

// VMTemplateService performs the get operations of the vm template
type VMTemplateService Service

// NewVMTemplateService creates a new vm template service
func NewVMTemplateService(doer doer, serviceURL string, stamp string) *VMTemplateService {
	url, err := url.Parse(serviceURL)
	if err != nil {
		return nil
	}
	return &VMTemplateService{doer: doer, serviceURL: *url, stampID: stamp}
}

type vmTemplateFeed struct {
	VMTemplates []*VMTemplate `json:"value"`
}

// VMTemplate object
type VMTemplate struct {
	StampID                                *string          `json:"StampId"`
	ID                                     *string          `json:"ID"`
	AccessedTime                           *string          `json:"AccessedTime"`
	AddedTime                              *string          `json:"AddedTime"`
	Admin                                  *string          `json:"Admin"`
	AdminPasswordHasValue                  *bool            `json:"AdminPasswordHasValue"`
	ComputerName                           *string          `json:"ComputerName"`
	CPUCount                               *uint8           `json:"CPUCount"`
	CPUMax                                 *int32           `json:"CPUMax"`
	CPUReserve                             *int32           `json:"CPUReserve"`
	CPUType                                *string          `json:"CPUType"`
	CreationTime                           *string          `json:"CreationTime"`
	DiskIO                                 *int32           `json:"DiskIO"`
	DomainAdmin                            *string          `json:"DomainAdmin"`
	DomainAdminPasswordHasValue            *bool            `json:"DomainAdminPasswordHasValue"`
	ExpectedCPUUtilization                 *int32           `json:"ExpectedCPUUtilization"`
	Enabled                                *bool            `json:"Enabled"`
	FullName                               *string          `json:"FullName"`
	HasVMAdditions                         *bool            `json:"HasVMAdditions"`
	IsHighlyAvailable                      *bool            `json:"IsHighlyAvailable"`
	JoinDomain                             *string          `json:"JoinDomain"`
	JoinWorkgroup                          *string          `json:"JoinWorkgroup"`
	LibraryGroup                           *string          `json:"LibraryGroup"`
	LimitCPUForMigration                   *bool            `json:"LimitCPUForMigration"`
	LimitCPUFunctionality                  *bool            `json:"LimitCPUFunctionality"`
	Location                               *string          `json:"Location"`
	Memory                                 *int32           `json:"Memory"`
	MergeAnswerFile                        *bool            `json:"MergeAnswerFile"`
	ModifiedTime                           *string          `json:"ModifiedTime"`
	Name                                   *string          `json:"Name"`
	NetworkUtilization                     *int32           `json:"NetworkUtilization"`
	OperatingSystem                        *string          `json:"OperatingSystem"`
	OperatingSystemInstance                *OperatingSystem `json:"OperatingSystemInstance"`
	OSType                                 *string          `json:"OSType"`
	OrgName                                *string          `json:"OrgName"`
	Owner                                  *UserAndRole     `json:"Owner"`
	GrantedToList                          []*UserAndRole   `json:"GrantedToList"`
	QuotaPoint                             *int32           `json:"QuotaPoint"`
	ProductKeyHasValue                     *bool            `json:"ProductKeyHasValue"`
	RelativeWeight                         *int32           `json:"RelativeWeight"`
	ShareSCSIBus                           *bool            `json:"ShareSCSIBus"`
	Shielded                               *bool            `json:"Shielded"`
	Tag                                    *string          `json:"Tag"`
	TimeZone                               *int32           `json:"TimeZone"`
	TotalVHDCapacity                       *string          `json:"TotalVHDCapacity"`
	UndoDisksEnabled                       *bool            `json:"UndoDisksEnabled"`
	UseHardwareAssistedVirtualization      *bool            `json:"UseHardwareAssistedVirtualization"`
	Accessibility                          *string          `json:"Accessibility"`
	CostCenter                             *string          `json:"CostCenter"`
	Description                            *string          `json:"Description"`
	IsTagEmpty                             *bool            `json:"IsTagEmpty"`
	NicCount                               *int32           `json:"NicCount"`
	NumLockEnabled                         *bool            `json:"NumLockEnabled"`
	VMAddition                             *string          `json:"VMAddition"`
	IsCustomizable                         *bool            `json:"IsCustomizable"`
	DomainAdminPasswordIsServiceSetting    *bool            `json:"DomainAdminPasswordIsServiceSetting"`
	SANCopyCapable                         *bool            `json:"SANCopyCapable"`
	IsTemporaryTemplate                    *bool            `json:"IsTemporaryTemplate"`
	VMTemplateID                           *string          `json:"VMTemplateId"`
	VirtualHardDiskID                      *string          `json:"VirtualHardDiskId"`
	VMId                                   *string          `json:"VMId"`
	SharePath                              *string          `json:"SharePath"`
	ApplicationProfileID                   *string          `json:"ApplicationProfileId"`
	CloudID                                *string          `json:"CloudId"`
	DynamicMemoryBufferPercentage          *int16           `json:"DynamicMemoryBufferPercentage"`
	DynamicMemoryEnabled                   *bool            `json:"DynamicMemoryEnabled"`
	DynamicMemoryMaximumMB                 *int32           `json:"DynamicMemoryMaximumMB"`
	MemoryWeight                           *int16           `json:"MemoryWeight"`
	DynamicMemoryPreferredBufferPercentage *int16           `json:"DynamicMemoryPreferredBufferPercentage"`
	SQLProfileID                           *string          `json:"SQLProfileId"`
	VirtualFloppyDriveID                   *string          `json:"VirtualFloppyDriveId"`
	BootOrder                              []*string        `json:"BootOrder"`
	CustomProperties                       []*string        `json:"CustomProperties"`
	GuiRunOnceCommands                     []*string        `json:"GuiRunOnceCommands"`
	ServerFeatures                         []*string        `json:"ServerFeatures"`
	Status                                 *string          `json:"Status"`
	VirtualizationPlatform                 *string          `json:"VirtualizationPlatform"`
	CapabilityProfile                      *string          `json:"CapabilityProfile"`
	AutoLogonCount                         *int32           `json:"AutoLogonCount"`
	DomainJoinOrganizationalUnit           *string          `json:"DomainJoinOrganizationalUnit"`
	SANStatus                              []*string        `json:"SANStatus"`
	Generation                             *int32           `json:"Generation"`
}

// GetByID gets a template by its ID
func (s VMTemplateService) GetByID(ID string) (*VMTemplate, error) {
	template := &VMTemplate{}
	uri := fmt.Sprintf(s.serviceURL.String()+"(ID=guid'%s',StampId=guid'%s')", ID, s.stampID)
	body, err := get(s, uri)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, template)
	if err != nil {
		return nil, err
	}
	return template, nil
}

// GetAll get all the vm template objects
func (s VMTemplateService) GetAll() ([]*VMTemplate, error) {
	var VMTemplates []*VMTemplate
	var feed vmTemplateFeed
	body, err := get(s, s.serviceURL.String())
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &feed)
	if err != nil {
		return nil, err
	}
	for _, template := range feed.VMTemplates {
		VMTemplates = append(VMTemplates, template)
	}
	return VMTemplates, nil
}
