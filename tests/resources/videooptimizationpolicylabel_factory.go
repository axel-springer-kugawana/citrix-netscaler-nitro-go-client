package resources

import (
	"testing"

	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
)

func Setup_videooptimizationpolicylabel(t *testing.T, client *nitro.NitroClient) (*nitro.Videooptimizationpolicylabel, func()) {
	//Je pense pas qu'il est testable
	// resource := nitro.Videooptimizationpolicylabel{
	// Labelname:       "videooptimizationpolicylabel",
	// Comment:         "test",
	// Policylabeltype: "videoopt_res",
	// }

	return nil, func() {
		//client.DeleteVideooptimizationpolicylabel(resource.ToKey())
	}
}
