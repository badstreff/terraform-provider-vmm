package getting

type virtualNetworkAdapterFeed struct {
	VirtualNetworkAdapters []*VirtualNetworkAdapter `json:"value"`
}

// VirtualNetworkAdapter object
type VirtualNetworkAdapter struct {
	ID                           *string
	Name                         *string
	VMwarePortGroup              *string
	EthernetAddressType          *string
	PhysicalAddressType          *string
	EthernetAddress              *string
	PhysicalAddress              *string
	VirtualNetworkAdapterType    *string
	Location                     *string
	Tag                          *string
	Description                  *string
	VMNetworkID                  *string `json:"VMNetworkId"`
	VMNetworkName                *string
	VMSubnetID                   *string `json:"VMSubnetId"`
	TemplateID                   *string `json:"TemplateId"`
	VMId                         *string
	StampID                      *string `json:"StampId"`
	MACAddressesSpoofingEnabled  *bool
	SlotID                       *int `json:"SlotId"`
	VLanEnabled                  *bool
	VLanID                       *int16 `json:"VLanId"`
	VMNetworkOptimizationEnabled *bool
	VmwAdapterIndex              *int
	Accessibility                *string
	AddedTime                    *string
	ChildObjectIDs               []*string
	Enabled                      *bool
	IPv4AddressType              *string
	IPv6AddressType              *string
	IPv4Addresses                []*string
	IPv6Addresses                []*string
	IPv4AddressPoolsID           []*string `json:"IPv4AddressPoolsId"`
	IPv6AddressPoolsID           []*string `json:"IPv6AddressPoolsId"`
	MACAddress                   *string
	MACAddressesSpoolingEnabled  *bool
	MACAddressSpoofingEnabled    *bool
	EnableMACAddressSpoofing     *bool
	MACAddressType               *string
	ModifiedTime                 *string
	ParentID                     *string `json:"ParentId"`
	RequiredBandwidth            *string
	IsSynthetic                  *bool
}
