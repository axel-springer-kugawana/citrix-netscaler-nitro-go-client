package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransformaction(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_transformaction()

	if resource == nil {
		return
	}

	err := client.AddTransformaction(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsTransformaction(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err := client.ExistsTransformaction(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetTransformaction(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListTransformaction()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateTransformaction(*resource)

	// assert.NoError(t, err)

	err = client.RenameTransformaction(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteTransformaction(resource.Name + "-rename")

	assert.NoError(t, err)
}
