package client

import (
	"github.com/badstreff/terraform-provider-vmm/pkg/spf-vmm-api/creating"
	"github.com/badstreff/terraform-provider-vmm/pkg/spf-vmm-api/deleting"
	"github.com/badstreff/terraform-provider-vmm/pkg/spf-vmm-api/getting"
)

// VirtualDiskDriveService service
type VirtualDiskDriveService struct {
	creating.VirtualDiskDriveCreator
	deleting.GenericDeleter
	getting.VirtualDiskDriveGetter
}
