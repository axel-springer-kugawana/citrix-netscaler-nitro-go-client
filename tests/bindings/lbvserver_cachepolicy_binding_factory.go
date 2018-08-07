package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func setup_lbvserver_cachepolicy_binding(t *testing.T, client *nitro.NitroClient) (*nitro.LbvserverCachepolicyBinding, func()) {
	//resource := nitro.LbvserverCachepolicyBinding {
	//Bindpoint: (REQUEST|RESPONSE),
	//Gotopriorityexpression: string,
	//Invoke: bool,
	//Labelname: string,
	//Labeltype: (reqvserver|resvserver|policylabel),
	//Name: lbvserver.name,
	//Policyname: cachepolicy.policyname,
	//Priority: double,
	//}

	return nil, func() {
	}
}
