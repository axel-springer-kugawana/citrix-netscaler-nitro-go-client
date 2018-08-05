package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNetprofile(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_netprofile()

	if resource == nil {
		return
	}

	err := client.AddNetprofile(*resource)

	assert.NoError(t, err)

	res, err := client.GetNetprofile(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListNetprofile()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateNetprofile(*resource)

	// assert.NoError(t, err)

	err = client.RenameNetprofile(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteNetprofile(resource.Name + "-rename")

	assert.NoError(t, err)
}
