package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func Setup_lbgroup(t *testing.T, client *nitro.NitroClient) (*nitro.Lbgroup, func()) {
	//resource := nitro.Lbgroup {
	//Name: "lbgroup",
	//Backuppersistencetimeout: double,
	//Cookiedomain: string,
	//Cookiename: string,
	//Persistencebackup: (SOURCEIP|NONE),
	//Persistencetype: (SOURCEIP|COOKIEINSERT|RULE|NONE),
	//Persistmask: string,
	//Rule: string,
	//Timeout: double,
	//Usevserverpersistency: (ENABLED|DISABLED),
	//V6persistmasklen: double,
	//}

	return nil, func() {
	}
}
