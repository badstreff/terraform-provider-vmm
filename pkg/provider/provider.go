package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider the entry point for the provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Base URL for the SPF endpoint",
				DefaultFunc: nil,
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Username for SPF operations",
				DefaultFunc: nil,
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Password for the user account",
				DefaultFunc: nil,
			},
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Domain for the user account",
				DefaultFunc: nil,
			},
			"stamp_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Stamp ID to use for all SPF operations",
				DefaultFunc: nil,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"vmm_virtual_disk_drive": resourceVirtualDiskDrive(),
			"vmm_virtual_machine":    resourceVirtualMachine(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		BaseURL:  d.Get("base_url").(string),
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		Domain:   d.Get("domain").(string),
		StampID:  d.Get("stamp_id").(string),
	}
	return config.NewClient()
}
