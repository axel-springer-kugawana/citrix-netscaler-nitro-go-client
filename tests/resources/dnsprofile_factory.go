package resources

import (
	"testing"

	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
)

func Setup_dnsprofile(t *testing.T, client *nitro.NitroClient) (*nitro.Dnsprofile, func()) {
	resource := nitro.Dnsprofile{
		Dnsprofilename: "dnsprofile",
		//Cacheecsresponses:      "DISABLED",
		Cachenegativeresponses: "DISABLED",
		Cacherecords:           "DISABLED",
		Dnsanswerseclogging:    "ENABLED",
		Dnserrorlogging:        "ENABLED",
		Dnsextendedlogging:     "ENABLED",
		Dnsquerylogging:        "ENABLED",
		Dropmultiqueryrequest:  "DISABLED",
	}

	return &resource, func() {
	}
}
