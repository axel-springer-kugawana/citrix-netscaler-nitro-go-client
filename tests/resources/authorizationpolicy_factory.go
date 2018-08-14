package resources

import (
	"testing"

	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
)

func Setup_authorizationpolicy(t *testing.T, client *nitro.NitroClient) (*nitro.Authorizationpolicy, func()) {
	resource := nitro.Authorizationpolicy{
		Name:   "authorizationpolicy",
		Action: "ALLOW",
		Rule:   "is_vpn_url",
	}

	return &resource, func() {
		client.DeleteAuthorizationpolicy(resource.ToKey())
	}
}
