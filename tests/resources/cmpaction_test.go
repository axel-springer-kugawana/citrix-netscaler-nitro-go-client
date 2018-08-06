package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCmpaction(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := setup_cmpaction(t)

	if resource == nil {
		return
	}

	err := client.AddCmpaction(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsCmpaction(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsCmpaction(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetCmpaction(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListCmpaction()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateCmpaction(*resource)

	// assert.NoError(t, err)

	err = client.RenameCmpaction(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteCmpaction(resource.Name + "-rename")

	assert.NoError(t, err)
}
