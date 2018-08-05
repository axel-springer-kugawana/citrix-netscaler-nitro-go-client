package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
)

func create_appflowcollector() *nitro.Appflowcollector {
	resource := nitro.Appflowcollector {
		Name: "appflowcollector",
		Ipaddress: "1.3.5.7",
		// Netprofile: netprofile.name,
		// Port: int,
		// Transport: (ipfix|logstream),
	}
	return &resource
}
