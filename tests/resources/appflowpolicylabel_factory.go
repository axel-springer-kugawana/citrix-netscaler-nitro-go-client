package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func setup_appflowpolicylabel(t *testing.T, client *nitro.NitroClient) (*nitro.Appflowpolicylabel, func()) {
	resource := nitro.Appflowpolicylabel{
		Labelname:       "appflowpolicylabel",
		Policylabeltype: "HTTP",
	}

	return &resource, func() {
	}
}
