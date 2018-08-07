package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func Setup_rewritepolicylabel(t *testing.T, client *nitro.NitroClient) (*nitro.Rewritepolicylabel, func()) {
	//resource := nitro.Rewritepolicylabel {
	//Labelname: "rewritepolicylabel",
	//Comment: string,
	//Transform: (http_req|http_res|othertcp_req|othertcp_res|url|text|clientless_vpn_req|clientless_vpn_res|sipudp_req|sipudp_res|siptcp_req|siptcp_res|diameter_req|diameter_res|radius_req|radius_res|dns_req|dns_res),
	//}

	return nil, func() {
	}
}
