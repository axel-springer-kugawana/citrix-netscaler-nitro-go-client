package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func Setup_responderaction(t *testing.T, client *nitro.NitroClient) (*nitro.Responderaction, func()) {
	//resource := nitro.Responderaction {
	//Name: "responderaction",
	//Bypasssafetycheck: (YES|NO),
	//Comment: string,
	//Htmlpage: string,
	//Reasonphrase: string,
	//Responsestatuscode: double,
	//Target: string,
	//Type: (noop|respondwith|redirect|respondwithhtmlpage|sqlresponse_ok|sqlresponse_error),
	//}

	return nil, func() {
	}
}
