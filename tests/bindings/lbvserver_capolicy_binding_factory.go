package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func Setup_lbvserver_capolicy_binding(t *testing.T, client *nitro.NitroClient) (*nitro.LbvserverCapolicyBinding, func()) {
	//resource := nitro.LbvserverCapolicyBinding {
	//Bindpoint: (REQUEST|RESPONSE),
	//Gotopriorityexpression: string,
	//Invoke: bool,
	//Labelname: string,
	//Labeltype: (reqvserver|resvserver|policylabel),
	//Name: lbvserver.name,
	//Policyname: capolicy.name,
	//Priority: double,
	//}

	return nil, func() {
	}
}
