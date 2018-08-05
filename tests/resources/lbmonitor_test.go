package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLbmonitor(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_lbmonitor()

	if resource == nil {
		return
	}

	err := client.AddLbmonitor(*resource)

	assert.NoError(t, err)

}
