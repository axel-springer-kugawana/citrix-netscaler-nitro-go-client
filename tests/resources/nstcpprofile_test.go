package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNstcpprofile(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_nstcpprofile()

	if resource == nil {
		return
	}

	err := client.AddNstcpprofile(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsNstcpprofile(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err := client.ExistsNstcpprofile(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetNstcpprofile(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListNstcpprofile()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateNstcpprofile(*resource)

	// assert.NoError(t, err)

	err = client.RenameNstcpprofile(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteNstcpprofile(resource.Name + "-rename")

	assert.NoError(t, err)
}
