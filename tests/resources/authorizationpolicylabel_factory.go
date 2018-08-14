package resources

import (
	"testing"

	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
)

func Setup_authorizationpolicylabel(t *testing.T, client *nitro.NitroClient) (*nitro.Authorizationpolicylabel, func()) {
	// resource := nitro.Authorizationpolicylabel{
	// 	Labelname: "authorizationpolicylabel",
	// }

	return nil, func() {
		//client.DeleteAuthorizationpolicylabel(resource.ToKey())
	}
}
