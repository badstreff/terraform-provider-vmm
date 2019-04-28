package provider

import (
	"github.com/badstreff/terraform-provider-vmm/pkg/spf-vmm-api/client"
	"github.com/badstreff/terraform-provider-vmm/pkg/spf-vmm-api/creating"
	"github.com/hashicorp/terraform/helper/schema"
	"time"
)

func resourceVirtualDiskDrive() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreateVirtualDiskDrive,
		Read:   resourceReadVirtualDiskDrive,
		Delete: resourceDeleteVirtualDiskDrive,
		Schema: map[string]*schema.Schema{
			"vm_name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: nil,
			},
			"vm_id": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: nil,
			},
			"bus": &schema.Schema{
				Type:         schema.TypeInt,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: nil,
			},
			"file_name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: nil,
			},
			"ide": &schema.Schema{
				Type:         schema.TypeBool,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: nil,
			},
			"lun": &schema.Schema{
				Type:         schema.TypeInt,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: nil,
			},
			"path": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: nil,
			},
			"scsi": &schema.Schema{
				Type:         schema.TypeBool,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: nil,
			},
			"stamp_id": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: nil,
			},
			"hard_disk_name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: nil,
			},
			"hard_disk_id": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: nil,
			},
		},
	}
}

func resourceCreateVirtualDiskDrive(d *schema.ResourceData, meta interface{}) error {
	var vmid string
	var diskid string
	var ide *bool
	var scsi *bool
	client := meta.(*client.Client)
	bus := uint8(d.Get("bus").(int))
	lun := uint8(d.Get("lun").(int))
	stampID := d.Get("stamp_id").(string)
	fileName := d.Get("file_name").(string)

	if v := d.Get("scsi").(bool); v != false {
		scsi = &v
	}

	if v := d.Get("ide").(bool); v != false {
		ide = &v
	}

	// set vmid based on name or id passed
	if vmName := d.Get("vm_name").(string); vmName != "" {
		vm, _ := client.VirtualMachines.GetByName(vmName)
		vmid = *vm.ID
	} else {
		vmid = d.Get("vm_id").(string)
	}

	// set disk id based on name or id passed
	if diskName := d.Get("hard_disk_name").(string); diskName != "" {
		disk, _ := client.VirtualHardDisks.GetByName(diskName)
		diskid = *disk.ID
	} else {
		diskid = d.Get("hard_disk_id").(string)
	}

	s := func(s string) *string { return &s }
	drive := creating.VirtualDiskDrive{
		ODataType:         s("VMM.VirtualDiskDrive"),
		StampID:           &stampID,
		VMID:              &vmid,
		VirtualHardDiskID: &diskid,
		FileName:          &fileName,
		Bus:               &bus,
		LUN:               &lun,
		SCSI:              scsi,
		IDE:               ide,
	}
	id, err := client.VirtualDiskDrives.Create(&drive)
	if err != nil {
		return err
	}
	time.Sleep(2 * time.Second) // Give a second for the job to queue
	vm, _ := client.VirtualMachines.GetByID(vmid)
	client.Jobs.WaitForJobToComplete(*vm.MostRecentTaskID, 30)
	d.SetId(*id)
	return resourceReadVirtualDiskDrive(d, meta)
}

func resourceDeleteVirtualDiskDrive(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Client)
	err := client.VirtualDiskDrives.DeleteByID(d.Id())
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceReadVirtualDiskDrive(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Client)
	drive, err := client.VirtualDiskDrives.GetByID(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}
	d.Set("path", *drive.Path)
	d.Set("bus", *drive.Bus)
	d.Set("lun", *drive.LUN)
	d.Set("ide", *drive.IDE)
	d.Set("path", *drive.Path)
	d.Set("stamp_id", *drive.StampID)
	return nil
}
