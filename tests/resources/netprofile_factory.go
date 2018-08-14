package resources

import (
	"testing"

	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
)

func Setup_netprofile(t *testing.T, client *nitro.NitroClient) (*nitro.Netprofile, func()) {
	resource := nitro.Netprofile{
		Name:             "netprofile",
		Overridelsn:      "ENABLED",
		Srcip:            "10.2.0.144",
		Srcippersistency: "DISABLED",
		//	Td:               1,
	}

	return &resource, func() {
	}
}
