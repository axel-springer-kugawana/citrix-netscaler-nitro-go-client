package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func Setup_appflowcollector(t *testing.T, client *nitro.NitroClient) (*nitro.Appflowcollector, func()) {
	resource := nitro.Appflowcollector{
		Name:      "appflowcollector",
		Ipaddress: "1.3.5.7",
		// Netprofile: netprofile.name,
		Port: 1234,
	}

	return &resource, nil
}
