package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func setup_lbvserver_transformpolicy_binding(t *testing.T, client *nitro.NitroClient) (*nitro.LbvserverTransformpolicyBinding, func()) {
	//resource := nitro.LbvserverTransformpolicyBinding {
	//Bindpoint: (REQUEST|RESPONSE),
	//Gotopriorityexpression: string,
	//Invoke: bool,
	//Labelname: string,
	//Labeltype: (reqvserver|resvserver|policylabel),
	//Name: lbvserver.name,
	//Policyname: transformpolicy.name,
	//Priority: double,
	//}

	return nil, func() {
	}
}
