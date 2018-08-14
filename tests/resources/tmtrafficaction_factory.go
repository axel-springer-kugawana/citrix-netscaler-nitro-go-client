package resources

import (
	"testing"

	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
)

func Setup_tmtrafficaction(t *testing.T, client *nitro.NitroClient) (*nitro.Tmtrafficaction, func()) {
	// resource := nitro.Tmtrafficaction{
	// 	Name:             "tmtrafficaction",
	// 	Apptimeout:       60,
	// 	Forcedtimeout:    "RESET",
	// 	Forcedtimeoutval: 60,
	// 	//Formssoaction:    "POST",
	// 	Initiatelogout:   "OFF",
	// 	Kcdaccount:       "1",
	// 	Passwdexpression: "",
	// 	Persistentcookie: "OFF",
	// 	Samlssoprofile:   "",
	// 	Sso:              "OFF",
	// 	Userexpression:   "",
	// }

	return nil, func() {
		//client.DeleteTmtrafficaction(resource.ToKey())
	}
}
