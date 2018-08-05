package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
)

func create_server() *nitro.Server {
	resource := nitro.Server{
		Name:      "server",
		Ipaddress: "1.3.5.7",
		// Domain: "abc.com",
		Comment: "comment",
		State:   "ENABLED",
		// Ipv6address: "NO",
		// Td: 0,
		// Domainresolveretry: 5,
		// Translationip: "1.2.3.4",
		// Translationmask: "5.6.7.8",
		// Graceful: "YES",
		// Delay: 123,
		// Domainresolvenow: true,
	}

	return &resource
}
