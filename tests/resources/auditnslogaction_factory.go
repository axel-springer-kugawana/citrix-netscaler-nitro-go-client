package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func Setup_auditnslogaction(t *testing.T, client *nitro.NitroClient) (*nitro.Auditnslogaction, func()) {
	resource := nitro.Auditnslogaction{
		Name:                "auditnslogaction",
		Serverip:            "1.3.5.7",
		Serverport:          1234,
		Dateformat:          "MMDDYYYY",
		Logfacility:         "LOCAL0",
		Tcp:                 "ALL",
		Acl:                 "ENABLED",
		Timezone:            "GMT_TIME",
		Userdefinedauditlog: "NO",
		Appflowexport:       "ENABLED",
		Lsn:                 "ENABLED",
		Alg:                 "ENABLED",
	}

	resource.Loglevel = append(resource.Loglevel, "ALL")

	return &resource, func() {
	}
}
