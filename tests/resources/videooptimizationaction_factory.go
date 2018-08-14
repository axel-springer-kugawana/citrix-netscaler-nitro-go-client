package resources

import (
	"testing"

	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
)

func Setup_videooptimizationaction(t *testing.T, client *nitro.NitroClient) (*nitro.Videooptimizationaction, func()) {
	// resource := nitro.Videooptimizationaction{
	// 	Name:    "video",
	// 	Comment: "sdfsdfsd",
	// 	Rate:    2,
	// 	Type:    "clear_text_abr", // | clear_text_abr | encrypted_abr | trigger_enc_abr | optimize_abr)
	// }

	return nil, func() {
		//client.DeleteVideooptimizationaction(resource.ToKey())
	}
}
