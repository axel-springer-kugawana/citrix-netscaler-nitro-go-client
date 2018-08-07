package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func setup_lbmetrictable(t *testing.T, client *nitro.NitroClient) (*nitro.Lbmetrictable, func()) {
	resource := nitro.Lbmetrictable {
		Metrictable: "lbmetrictable",
	}

	return &resource, func() {
	}
}
