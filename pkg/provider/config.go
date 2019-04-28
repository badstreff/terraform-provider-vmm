package provider

import (
	"github.com/badstreff/terraform-provider-vmm/pkg/spf-vmm-api/client"
	"net/url"
)

// Config contains settings for creating our SPF client
type Config struct {
	BaseURL  string
	Username string
	Password string
	Domain   string
	StampID  string
}

// NewClient returns a new spf client for making api calls
func (c *Config) NewClient() (*client.Client, error) {
	credentials := client.NTLMCredentials{
		Username: c.Username,
		Password: c.Password,
		Domain:   c.Domain,
	}
	baseURL, err := url.Parse(c.BaseURL)
	client := client.NewClient(credentials, *baseURL, c.StampID)
	return client, err
}
