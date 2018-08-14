package resources

import (
	"testing"

	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
)

func Setup_dospolicy(t *testing.T, client *nitro.NitroClient) (*nitro.Dospolicy, func()) {
	resource := nitro.Dospolicy{
		Name:          "dospolicy",
		Cltdetectrate: 60,
		Qdepth:        21,
	}

	return &resource, func() {
		client.DeleteDospolicy(resource.ToKey())
	}
}
