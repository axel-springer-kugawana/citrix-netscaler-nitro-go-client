package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppqoeaction(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := setup_appqoeaction(t)

	if resource == nil {
		return
	}

	err := client.AddAppqoeaction(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsAppqoeaction(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsAppqoeaction(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetAppqoeaction(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListAppqoeaction()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateAppqoeaction(*resource)

	// assert.NoError(t, err)

	err = client.RenameAppqoeaction(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteAppqoeaction(resource.Name + "-rename")

	assert.NoError(t, err)
}
