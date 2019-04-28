package client

import (
	"crypto/tls"
	"github.com/badstreff/terraform-provider-vmm/pkg/spf-vmm-api/creating"
	"github.com/badstreff/terraform-provider-vmm/pkg/spf-vmm-api/deleting"
	"github.com/badstreff/terraform-provider-vmm/pkg/spf-vmm-api/getting"
	"github.com/badstreff/terraform-provider-vmm/pkg/spf-vmm-api/updating"
	"github.com/vadimi/go-http-ntlm"
	"net/http"
	"net/url"
)

// NTLMCredentials object
type NTLMCredentials struct {
	Username string
	Password string
	Domain   string
}

// Client used to make calls the the spf vmm api
type Client struct {
	baseURL           url.URL
	client            http.Client
	stampID           string
	Clouds            CloudService
	VirtualMachines   VirtualMachineService
	VirtualDiskDrives VirtualDiskDriveService
	VirtualHardDisks  VirtualHardDiskService
	VMTemplates       VMTemplateService
	Jobs              JobService
}

// NewClient constructs a client object for communicating with the Service Provider Foundation (SPF)
func NewClient(credentials NTLMCredentials, baseURL url.URL, stampID string) *Client {
	httpClient := http.Client{
		Transport: &httpntlm.NtlmTransport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Domain:          credentials.Domain,
			User:            credentials.Username,
			Password:        credentials.Password,
		},
	}
	client := Client{baseURL: baseURL, client: httpClient, stampID: stampID}
	client.Clouds = CloudService{
		getting.NewCloudService(&httpClient, baseURL.String()+"/Clouds", stampID),
	}
	client.VirtualHardDisks = VirtualHardDiskService{
		getting.NewVirtualHardDiskService(&httpClient, baseURL.String()+"/VirtualHardDisks", stampID),
	}
	client.VirtualMachines = VirtualMachineService{
		creating.NewVirtualMachineService(&httpClient, baseURL.String()+"/VirtualMachines", stampID),
		deleting.NewGenericDeleterService(&httpClient, baseURL.String()+"/VirtualMachines", stampID),
		getting.NewVirtualMachineService(&httpClient, baseURL.String()+"/VirtualMachines", stampID),
		updating.NewVirtualMachineService(&httpClient, baseURL.String()+"/VirtualMachines", stampID),
	}
	client.VirtualDiskDrives = VirtualDiskDriveService{
		creating.NewVirtualDiskDriveService(&httpClient, baseURL.String()+"/VirtualDiskDrives", stampID),
		deleting.NewGenericDeleterService(&httpClient, baseURL.String()+"/VirtualDiskDrives", stampID),
		getting.NewVirtualDiskDriveService(&httpClient, baseURL.String()+"/VirtualDiskDrives", stampID),
	}
	client.Jobs = JobService{
		getting.NewJobService(&httpClient, baseURL.String()+"/Jobs", stampID),
		updating.NewJobService(&httpClient, baseURL.String()+"/Jobs", stampID),
	}
	client.VMTemplates = VMTemplateService{
		getting.NewVMTemplateService(&httpClient, baseURL.String()+"/VMTemplates", stampID),
	}
	return &client
}
