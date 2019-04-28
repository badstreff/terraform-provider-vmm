package provider

import (
	"github.com/badstreff/terraform-provider-vmm/pkg/spf-vmm-api/client"
	"github.com/badstreff/terraform-provider-vmm/pkg/spf-vmm-api/creating"
	"github.com/badstreff/terraform-provider-vmm/pkg/spf-vmm-api/updating"
	"github.com/hashicorp/terraform/helper/schema"
	"time"
)

func resourceVirtualMachine() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreateVirtualMachine,
		Read:   resourceReadVirtualMachine,
		Update: resourceUpdateVirtualMachine,
		Delete: resourceDeleteVirtualMachine,
		Schema: map[string]*schema.Schema{
			"vm_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cloud_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"cloud_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"stamp_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"template_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"template_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"cpu_count": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     false,
				ValidateFunc: nil,
			},
			"memory": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     false,
				ValidateFunc: nil,
			},
			"network_adapters": &schema.Schema{
				Type:         schema.TypeList,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: nil,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"subnet_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"mac_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"ipv4_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"ipv6_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"vlan_enabled": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"vlan_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourceCreateVirtualMachine(d *schema.ResourceData, meta interface{}) error {
	var cloudid string
	var templateid string

	client := meta.(*client.Client)
	vmName := d.Get("vm_name").(string)
	stampID := d.Get("stamp_id").(string)

	// set cloud id based on name or id passed
	if cloudName := d.Get("cloud_name").(string); cloudName != "" {
		cloud, _ := client.Clouds.GetByName(cloudName)
		cloudid = *cloud.ID
	} else {
		cloudid = d.Get("cloud_id").(string)
	}

	// set template id based on name or id passed
	if templateName := d.Get("template_name").(string); templateName != "" {
		template, _ := client.VMTemplates.GetByName(templateName)
		templateid = *template.ID
	} else {
		templateid = d.Get("template_id").(string)
	}

	cpuCount := evaluateUInt8Ptr(d, "cpu_count")
	memory := evaluateIntPtr(d, "memory")

	s := func(s string) *string {
		if s != "" {
			return &s
		}
		return nil
	}
	b := func(b bool) *bool { return &b }
	adapterInput := make([]*creating.NewVMVirtualNetworkAdapterInput, 0)
	for _, item := range d.Get("network_adapters").([]interface{}) {
		adapter := item.(map[string]interface{})
		adapterInput = append(adapterInput, &creating.NewVMVirtualNetworkAdapterInput{
			VMNetworkName:   s(adapter["network_name"].(string)),
			VMSubnetName:    s(adapter["subnet_name"].(string)),
			MACAddressType:  s(adapter["mac_type"].(string)),
			MACAddress:      s(adapter["mac"].(string)),
			IPv4AddressType: s(adapter["ipv4_type"].(string)),
			IPv6AddressType: s(adapter["ipv6_type"].(string)),
			VLanEnabled:     b(adapter["vlan_enabled"].(bool)),
		})
	}
	vm := creating.VirtualMachine{
		ODataType:                     s("VMM.VirtualMachine"),
		StampID:                       &stampID,
		Name:                          &vmName,
		CloudID:                       &cloudid,
		VMTemplateID:                  &templateid,
		NewVirtualNetworkAdapterInput: adapterInput,
		CPUCount:                      cpuCount,
		Memory:                        memory,
	}
	id, err := client.VirtualMachines.Create(&vm)
	if err != nil {
		return err
	}
	// client.VirtualMachines.Start(*id)
	time.Sleep(2 * time.Second) // Give a second for the job to queue
	res, _ := client.VirtualMachines.GetByID(*id)
	client.Jobs.WaitForJobToComplete(*res.MostRecentTaskID, 30)
	d.SetId(*id)
	return resourceReadVirtualMachine(d, meta)
}

func resourceReadVirtualMachine(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Client)
	vm, err := client.VirtualMachines.GetByID(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}
	d.Set("vm_name", *vm.Name)
	d.Set("stamp_id", *vm.StampID)
	d.Set("cloud_id", *vm.CloudID)

	// TODO cleanup and roll into its own function
	// this is also kinda hacky and we need to make sure
	// that order is always correct, all we are really
	// doing here is setting the mac address to what scvmm
	// generates
	s := func(s string) *string { return &s }
	b := func(b bool) *bool { return &b }
	adapterInput := make([]*creating.NewVMVirtualNetworkAdapterInput, 0)
	adapters, _ := client.VirtualMachines.GetVirtualNetworkAdapters(d.Id())
	for i, item := range d.Get("network_adapters").([]interface{}) {
		adapter := item.(map[string]interface{})
		adapterInput = append(adapterInput, &creating.NewVMVirtualNetworkAdapterInput{
			VMNetworkName:   s(adapter["network_name"].(string)),
			VMSubnetName:    s(adapter["subnet_name"].(string)),
			MACAddressType:  s(adapter["mac_type"].(string)),
			MACAddress:      s(adapter["mac"].(string)),
			IPv4AddressType: s(adapter["ipv4_type"].(string)),
			IPv6AddressType: s(adapter["ipv6_type"].(string)),
			VLanEnabled:     b(adapter["vlan_enabled"].(bool)),
		})
		adapterInput[i].MACAddress = adapters[i].MACAddress
	}
	d.Set("network_adapters", flattenNetworkAdapters(adapterInput))
	return nil
}

func resourceUpdateVirtualMachine(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Client)
	vmName := d.Get("vm_name").(string)
	stampID := d.Get("stamp_id").(string)
	cpuCount := evaluateUInt8Ptr(d, "cpu_count")
	memory := evaluateIntPtr(d, "memory")

	id := d.Id()

	res, _ := client.VirtualMachines.GetByID(id)
	initialState := *res.VirtualMachineState
	if initialState == "Running" {
		client.VirtualMachines.Stop(id)
	}

	s := func(s string) *string { return &s }
	vm := updating.VirtualMachine{
		ODataType: s("VMM.VirtualMachine"),
		StampID:   &stampID,
		Name:      &vmName,
		ID:        &id,
		CPUCount:  cpuCount,
		Memory:    memory,
	}
	err := client.VirtualMachines.Update(&vm)
	if err != nil {
		return err
	}
	time.Sleep(2 * time.Second) // Give a second for the job to queue
	res, _ = client.VirtualMachines.GetByID(id)
	client.Jobs.WaitForJobToComplete(*res.MostRecentTaskID, 30)
	if initialState == "Running" {
		client.VirtualMachines.Start(id)
	}
	return resourceReadVirtualMachine(d, meta)
}

func resourceDeleteVirtualMachine(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Client)
	err := client.VirtualMachines.Stop(d.Id())
	if err != nil {
		return err
	}
	err = client.VirtualMachines.DeleteByID(d.Id())
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func flattenNetworkAdapters(in []*creating.NewVMVirtualNetworkAdapterInput) []map[string]interface{} {
	var out = make([]map[string]interface{}, len(in), len(in))
	for i, v := range in {
		m := make(map[string]interface{})
		m["network_name"] = *v.VMNetworkName
		m["mac_type"] = *v.MACAddressType
		m["mac"] = *v.MACAddress
		out[i] = m
	}
	return out
}
