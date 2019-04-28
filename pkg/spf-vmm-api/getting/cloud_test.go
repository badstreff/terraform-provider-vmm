package getting

import (
	"reflect"
	"testing"
)

func TestCloudGetByID(t *testing.T) {
	doer := MockDoer{}
	doer.StatusCode = 200
	doer.Data = `{
    "Description": "Services considered core to the function and stability of the testlab.",
    "ID": "f3c2cba8-5913-44c1-ac25-659dfda472d4",
    "LastModifiedDate": null,
    "Name": "Production",
    "ShieldedVMSupportPolicy": "ShieldedVMNotSupported",
    "StampId": "cc49f6b9-8db5-4392-9628-addbc9a46af6",
    "UserRoleID": null,
    "WritableLibraryPath": "\\\\VMM.EXAMPLE.COM\\MSSCVMMLibrary\\Production_Cloud\\",
    "odata.metadata": "https://VMM.EXAMPLE.COM:8090/SC2012R2/VMM/Microsoft.Management.Odata.svc/$metadata#Clouds/@Element"
}`

	cloudService := NewCloudService(doer, "", "")
	cloud, err := cloudService.GetByID("")
	if err != nil {
		t.Errorf("Error attempting to parse production cloud")
		t.Errorf("%v", err)
	}
	expected := &Cloud{
		Description:             sPtr("Services considered core to the function and stability of the testlab."),
		ID:                      sPtr("f3c2cba8-5913-44c1-ac25-659dfda472d4"),
		Name:                    sPtr("Production"),
		ShieldedVMSupportPolicy: sPtr("ShieldedVMNotSupported"),
		StampID:                 sPtr("cc49f6b9-8db5-4392-9628-addbc9a46af6"),
		WritableLibraryPath:     sPtr("\\\\VMM.EXAMPLE.COM\\MSSCVMMLibrary\\Production_Cloud\\"),
	}
	if !reflect.DeepEqual(cloud, expected) {
		t.Errorf("Error clouds do not match")
	}
}

func TestCloudGetAll(t *testing.T) {
	doer := MockDoer{}
	doer.StatusCode = 200
	doer.Data = `{
    "odata.metadata": "https://VMM.EXAMPLE.COM:8090/SC2012R2/VMM/Microsoft.Management.Odata.svc/$metadata#Clouds",
    "value": [
        {
            "Description": "",
            "ID": "f3d6b21f-641a-4415-9f77-3873e0faca64",
            "LastModifiedDate": null,
            "Name": "Test",
            "ShieldedVMSupportPolicy": "ShieldedVMNotSupported",
            "StampId": "cc49f6b9-8db5-4392-9628-addbc9a46af6",
            "UserRoleID": null,
            "WritableLibraryPath": "\\\\VMM.EXAMPLE.COM\\MSSCVMMLibrary\\Test_Cloud\\"
        },
        {
            "Description": "Services considered core to the function and stability of the testlab.",
            "ID": "f3c2cba8-5913-44c1-ac25-659dfda472d4",
            "LastModifiedDate": null,
            "Name": "Production",
            "ShieldedVMSupportPolicy": "ShieldedVMNotSupported",
            "StampId": "cc49f6b9-8db5-4392-9628-addbc9a46af6",
            "UserRoleID": null,
            "WritableLibraryPath": "\\\\VMM.EXAMPLE.COM\\MSSCVMMLibrary\\Production_Cloud\\"
        }
    ]
}`

	cloudService := NewCloudService(doer, "", "")
	Clouds, err := cloudService.GetAll()
	if err != nil {
		t.Errorf("Error attempting to parse production cloud feed")
		t.Errorf("%v", err)
	}
	if len(Clouds) != 2 {
		t.Errorf("Wrong number of clouds returned: expected 2, got %s", string(len(Clouds)))
	}
}

func TestNewCloudService(t *testing.T) {
	doer := MockDoer{}
	doer.StatusCode = 200
	service := NewCloudService(doer, "!@#^&%^*!$&!)@(#*!+)_@#+()!+", "")
	if service != nil {
		t.Errorf("Cloud service reported success on invalid url")
	}
}
