package resources

import (
	"testing"

	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
)

//Firewall policy
func Setup_appfwpolicy(t *testing.T, client *nitro.NitroClient) (*nitro.Appfwpolicy, func()) {
	resource := nitro.Appfwpolicy{
		Name:    "appfwpolicy",
		Comment: "Comment from Setup_appfwpolicy",
		//Logaction:   "TRUE",
		Profilename: "APPFW_BYPASS",
		Rule:        "TRUE",
	}

	return &resource, func() {
	}
}
