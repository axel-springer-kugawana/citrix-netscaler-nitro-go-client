package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func Setup_csvserver_feopolicy_binding(t *testing.T, client *nitro.NitroClient) (*nitro.CsvserverFeopolicyBinding, func()) {
	//resource := nitro.CsvserverFeopolicyBinding {
	//Bindpoint: (REQUEST|RESPONSE),
	//Gotopriorityexpression: string,
	//Invoke: bool,
	//Labelname: string,
	//Labeltype: (reqvserver|resvserver|policylabel),
	//Name: csvserver.name,
	//Policyname: spilloverpolicy.name,
	//Priority: double,
	//Targetlbvserver: lbvserver.name,
	//}

	return nil, func() {
	}
}
