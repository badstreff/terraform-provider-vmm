package getting

//import (
//	"reflect"
//	"testing"
//)
//
//func TestJobGetByID(t *testing.T) {
//	doer := MockDoer{}
//	doer.Data = `{
//    "AdditionalMessages": [
//        {
//            "CloudProblem": "",
//            "Code": "LibFailedToImportFileInUse",
//            "DetailedCode": 0,
//            "DetailedErrorCode": "",
//            "DetailedSource": "None",
//            "DisplayableErrorCode": "10803",
//            "ErrorCodeString": "10803",
//            "ErrorType": "Warning",
//            "ExceptionDetails": null,
//            "IsConditionallyTerminating": false,
//            "IsDeploymentBlocker": false,
//            "IsMomAlert": false,
//            "IsSuccess": false,
//            "IsTerminating": false,
//            "MessageParameters": "System.Collections.Generic.Dictionary` + "`" + `2[System.String,System.String]",
//            "MomAlertSeverity": "Error",
//            "Problem": "Unable to refresh * because the file is in use by another process.",
//            "RecommendedAction": "Wait for the next automatic library refresh, or manually refresh the library share after the process completes. Please contact your system administrator with this error ID.",
//            "RecommendedActionCLI": "Wait for the next automatic library refresh, or manually refresh the library share after the process completes. Please contact your system administrator with this error ID.",
//            "ShowDetailedError": false
//        }
//    ],
//    "AreAuditRecordsAvailable": false,
//    "CmdletName": "(Refresh Library - System Job)",
//    "CurrentStepId": null,
//    "Description": {
//        "CanSkipLastFailedStep": false,
//        "Code": "LibRefresher",
//        "DescriptionCodeString": "90",
//        "IsRestartable": false,
//        "IsStoppable": false,
//        "Name": "Update library",
//        "RequiresCredentialsForRestart": false
//    },
//    "EndTime": "2019-04-30T12:42:41.213-05:00",
//    "ErrorInfo": {
//        "CloudProblem": "",
//        "Code": "Success",
//        "DetailedCode": 0,
//        "DetailedErrorCode": "",
//        "DetailedSource": "None",
//        "DisplayableErrorCode": "0",
//        "ErrorCodeString": "0",
//        "ErrorType": "Error",
//        "ExceptionDetails": "",
//        "IsConditionallyTerminating": false,
//        "IsDeploymentBlocker": false,
//        "IsMomAlert": false,
//        "IsSuccess": true,
//        "IsTerminating": false,
//        "MessageParameters": "System.Collections.Generic.Dictionary` + "`" + `2[System.String,System.String]",
//        "MomAlertSeverity": "Error",
//        "Problem": "",
//        "RecommendedAction": "",
//        "RecommendedActionCLI": "",
//        "ShowDetailedError": false
//    },
//    "ID": "b71b1001-f634-4253-a4fd-ee0ec4a95576",
//    "IsCompleted": true,
//    "IsRestartable": false,
//    "IsStoppable": false,
//    "IsVisible": true,
//    "Name": "Update library",
//    "Owner": "EXAMPLE.COM\\ExampleUser",
//    "OwnerSID": null,
//    "PROTipID": null,
//    "Password": null,
//    "Progress": "100 %",
//    "ProgressValue": 100,
//    "ResultName": "Object Deleted",
//    "ResultObjectID": "30b5b2f1-3372-4b8c-87ac-d78751fac15a",
//    "ResultObjectType": "LibraryServer",
//    "ResultObjectTypeName": "Library Server",
//    "RootStepId": null,
//    "SkipLastFailedStep": null,
//    "Source": "",
//    "StampId": "4d140754-e603-447f-8982-e19495759aee",
//    "StartTime": "2019-04-30T12:42:37.01-05:00",
//    "Status": "SucceedWithInfo",
//    "StatusString": "Completed w/ Info",
//    "StepsLoaded": false,
//    "Target": "",
//    "TargetObjectID": "30b5b2f1-3372-4b8c-87ac-d78751fac15a",
//    "TargetObjectType": "LibraryServer",
//    "UserName": null,
//    "WasNotifiedOfCancel": false,
//    "odata.metadata": "https://VMM.EXAMPLE.COM:8090/SC2012R2/VMM/Microsoft.Management.Odata.svc/$metadata#Jobs/@Element"
//}`
//	jobService := NewJobService(doer, "", "")
//	job, err := jobService.GetByID("")
//	if err != nil {
//		t.Errorf("Error attempting to parse job")
//		t.Errorf("%v", err)
//	}
//	expected := &Job{}
//	if !reflect.DeepEqual(job, expected) {
//		t.Errorf("jobs not equal")
//	}
//}
//
//func TestJobGetAll(t *testing.T) {
//	doer := MockDoer{}
//	doer.Data = `
//`
//	jobService := NewJobService(doer, "", "")
//	Jobs, err := jobService.GetAll()
//	if err != nil {
//		t.Errorf("Error attempting to parse job feed")
//		t.Errorf("%v", err)
//	}
//	if len(Jobs) != 2 {
//		t.Errorf("Wrong number of jobs returned: expected 2, got %v", len(Jobs))
//	}
//}
//