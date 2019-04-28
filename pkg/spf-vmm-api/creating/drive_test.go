package creating

import (
	"testing"
)

func TestCreateDrive(t *testing.T) {
	doer := MockDoer{}
	doer.StatusCode = 200
	doer.Data = `<?xml version="1.0" encoding="utf-8"?>
<entry xml:base="https://example.com:8090/SC2012R2/VMM/Microsoft.Management.Odata.svc/" 
    xmlns="http://www.w3.org/2005/Atom" 
    xmlns:d="http://schemas.microsoft.com/ado/2007/08/dataservices" 
    xmlns:m="http://schemas.microsoft.com/ado/2007/08/dataservices/metadata">
    <id>https://example.com:8090/SC2012R2/VMM/Microsoft.Management.Odata.svc/VirtualDiskDrives(ID=guid'EXAMPLE-GUID-0001',StampId=guid'EXAMPLE-GUID-0002')</id>
    <category term="VMM.VirtualDiskDrive" scheme="http://schemas.microsoft.com/ado/2007/08/dataservices/scheme" />
    <link rel="edit" title="VirtualDiskDrive" href="VirtualDiskDrives(ID=guid'EXAMPLE-GUID-0001',StampId=guid'EXAMPLE-GUID-0002')" />
    <link rel="http://schemas.microsoft.com/ado/2007/08/dataservices/related/VirtualHardDisk" type="application/atom+xml;type=entry" title="VirtualHardDisk" href="VirtualDiskDrives(ID=guid'EXAMPLE-GUID-0001',StampId=guid'EXAMPLE-GUID-0002')/VirtualHardDisk" />
    <title />
    <updated>2019-04-19T15:41:22Z</updated>
    <author>
        <name />
    </author>
    <m:action metadata="https://example.com:8090/SC2012R2/VMM/Microsoft.Management.Odata.svc/$metadata#VMM.Expand" title="Expand" target="https://example.com:8090/SC2012R2/VMM/Microsoft.Management.Odata.svc/VirtualDiskDrives(ID=guid'EXAMPLE-GUID-0001',StampId=guid'EXAMPLE-GUID-0002')/Expand" />
    <content type="application/xml">
        <m:properties>
            <d:StampId m:type="Edm.Guid">EXAMPLE-GUID-0002</d:StampId>
            <d:ID m:type="Edm.Guid">EXAMPLE-GUID-0001</d:ID>
            <d:Bus m:type="Edm.Byte">0</d:Bus>
            <d:BusType>SCSI</d:BusType>
            <d:IsVHD m:type="Edm.Boolean">true</d:IsVHD>
            <d:LUN m:type="Edm.Byte">4</d:LUN>
            <d:Name>MyTestVM</d:Name>
            <d:VMId m:type="Edm.Guid" m:null="true" />
            <d:TemplateId m:type="Edm.Guid" m:null="true" />
            <d:ISOId m:type="Edm.Guid" m:null="true" />
            <d:HostDrive m:null="true" />
            <d:ISOLinked m:type="Edm.Boolean" m:null="true" />
            <d:Accessibility>Public</d:Accessibility>
            <d:Description></d:Description>
            <d:AddedTime m:type="Edm.DateTime">2019-04-12T17:04:14.967-05:00</d:AddedTime>
            <d:ModifiedTime m:type="Edm.DateTime">2019-04-19T00:20:07.381567-05:00</d:ModifiedTime>
            <d:Enabled m:type="Edm.Boolean">true</d:Enabled>
            <d:VirtualHardDiskId m:type="Edm.Guid">EXAMPLE-GUID-0003</d:VirtualHardDiskId>
            <d:VolumeType>None</d:VolumeType>
            <d:IDE m:type="Edm.Boolean" m:null="true" />
            <d:SCSI m:type="Edm.Boolean" m:null="true" />
            <d:FileName m:null="true" />
            <d:Path m:null="true" />
            <d:Size m:type="Edm.Int64">4194304</d:Size>
        </m:properties>
    </content>
</entry>`

	driveService := NewVirtualDiskDriveService(doer, "", "")
	expectedID := "EXAMPLE-GUID-0001"
	id, err := driveService.Create(&VirtualDiskDrive{})
	if err != nil {
		t.Errorf("Error creating drive")
		t.Errorf("%v", err)
	}

	if *id != expectedID {
		t.Errorf("Incorrect id, expected: %s, got %s.", expectedID, *id)
	}
}
