package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func setup_dnsprofile(t *testing.T, client *nitro.NitroClient) (*nitro.Dnsprofile, func()) {
	//resource := nitro.Dnsprofile {
	//Dnsprofilename: "dnsprofile",
	//Cacheecsresponses: (ENABLED|DISABLED),
	//Cachenegativeresponses: (ENABLED|DISABLED),
	//Cacherecords: (ENABLED|DISABLED),
	//Dnsanswerseclogging: (ENABLED|DISABLED),
	//Dnserrorlogging: (ENABLED|DISABLED),
	//Dnsextendedlogging: (ENABLED|DISABLED),
	//Dnsquerylogging: (ENABLED|DISABLED),
	//Dropmultiqueryrequest: (ENABLED|DISABLED),
	//}

	return nil, func() {
	}
}
