package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNshttpprofile(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_nshttpprofile()

	if resource == nil {
		return
	}

	err := client.AddNshttpprofile(*resource)

	assert.NoError(t, err)

	res, err := client.GetNshttpprofile(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListNshttpprofile()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateNshttpprofile(*resource)

	// assert.NoError(t, err)

	err = client.RenameNshttpprofile(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteNshttpprofile(resource.Name + "-rename")

	assert.NoError(t, err)
}
