package getting

//import (
//	"testing"
//)

//func TestDriveGetByID(t *testing.T) {
//	doer := MockDoer{}
//	doer.Data = ``
//	driveService := NewVirtualDiskDriveService(doer, "", "")
//	drive, err := driveService.GetByID("")
//	if err != nil {
//		t.Errorf("Error attempting to parse production drive")
//		t.Errorf("%v", err)
//	}
//	if *drive.Name != "" {
//		t.Errorf("Incorrect drive name, expected: %s, got %s.", "", *drive.Name)
//	}
//}

//func TestDriveGetAll(t *testing.T) {
//	doer := MockDoer{}
//	doer.Data = ``
//	driveService := NewVirtualDiskDriveService(doer, "", "")
//	drives, err := driveService.GetAll()
//	if err != nil {
//		t.Errorf("Error attempting to parse production drive feed")
//		t.Errorf("%v", err)
//	}
//	if len(drives) != 10 {
//		t.Errorf("Wrong number of drives returned: expected 10, got %v", len(drives))
//	}
//}
