package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDnsprofile(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_dnsprofile()

	if resource == nil {
		return
	}

	err := client.AddDnsprofile(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsDnsprofile(resource.Dnsprofilename + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err := client.ExistsDnsprofile(resource.Dnsprofilename)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetDnsprofile(resource.Dnsprofilename)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListDnsprofile()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateDnsprofile(*resource)

	// assert.NoError(t, err)

	err = client.RenameDnsprofile(resource.Dnsprofilename, resource.Dnsprofilename+"-rename")

	assert.NoError(t, err)

	err = client.DeleteDnsprofile(resource.Dnsprofilename + "-rename")

	assert.NoError(t, err)
}
