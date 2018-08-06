package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransformprofile(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := setup_transformprofile(t)

	if resource == nil {
		return
	}

	err := client.AddTransformprofile(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsTransformprofile(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsTransformprofile(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetTransformprofile(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListTransformprofile()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateTransformprofile(*resource)

	// assert.NoError(t, err)

	err = client.RenameTransformprofile(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteTransformprofile(resource.Name + "-rename")

	assert.NoError(t, err)
}
