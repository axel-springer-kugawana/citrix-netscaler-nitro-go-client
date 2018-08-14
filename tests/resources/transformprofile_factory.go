package resources

import (
	"testing"

	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
)

func Setup_transformprofile(t *testing.T, client *nitro.NitroClient) (*nitro.Transformprofile, func()) {
	resource := nitro.Transformprofile{
		Name: "transformprofile",
		//Comment: "",
		//Onlytransformabsurlinbody: "ON",
		Type: "URL",
	}

	return &resource, func() {
		client.DeleteTransformprofile(resource.ToKey())
	}
}
