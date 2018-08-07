package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func setup_lbvserver_cmppolicy_binding(t *testing.T, client *nitro.NitroClient) (*nitro.LbvserverCmppolicyBinding, func()) {
	//resource := nitro.LbvserverCmppolicyBinding {
	//Bindpoint: (REQUEST|RESPONSE),
	//Gotopriorityexpression: string,
	//Invoke: bool,
	//Labelname: string,
	//Labeltype: (reqvserver|resvserver|policylabel),
	//Name: lbvserver.name,
	//Policyname: cmppolicy.name,
	//Priority: double,
	//}

	return nil, func() {
	}
}
