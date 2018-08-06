package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func setup_appflowcollector(t *testing.T) *nitro.Appflowcollector {
	resource := nitro.Appflowcollector{
		Name:      "appflowcollector",
		Ipaddress: "1.3.5.7",
		// Netprofile: netprofile.name,
		Port: 1234,
	}

	return &resource
}
