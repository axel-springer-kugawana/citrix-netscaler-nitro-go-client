package resources

import (
	"testing"

	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
)

func Setup_tmsessionaction(t *testing.T, client *nitro.NitroClient) (*nitro.Tmsessionaction, func()) {
	// resource := nitro.Tmsessionaction{
	// 	Name: "tmsessionaction",
	// 	Defaultauthorizationaction: "ALLOW",
	// 	Homepage:                   "index",
	// 	Httponlycookie:             "YES",
	// 	Kcdaccount:                 "2",
	// 	//Persistentcookie:         "NO",
	// 	Persistentcookievalidity: 60,
	// 	Sesstimeout:              120,
	// 	Sso:                      "OFF",
	// 	Ssocredential:            "PRIMARY",
	// 	Ssodomain:                "sso.domain.com",
	// }

	return nil, func() {
		//client.DeleteTmsessionaction(resource.ToKey())
	}
}
