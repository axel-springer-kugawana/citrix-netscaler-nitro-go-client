package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDnsaction64(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := setup_dnsaction64(t)

	if resource == nil {
		return
	}

	err := client.AddDnsaction64(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsDnsaction64(resource.Actionname + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsDnsaction64(resource.Actionname)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetDnsaction64(resource.Actionname)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListDnsaction64()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateDnsaction64(*resource)

	// assert.NoError(t, err)

	err = client.RenameDnsaction64(resource.Actionname, resource.Actionname+"-rename")

	assert.NoError(t, err)

	err = client.DeleteDnsaction64(resource.Actionname + "-rename")

	assert.NoError(t, err)
}
