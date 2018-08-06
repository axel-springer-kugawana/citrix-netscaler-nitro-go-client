package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppflowaction(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_appflowaction()

	if resource == nil {
		return
	}

	err := client.AddAppflowaction(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsAppflowaction(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err := client.ExistsAppflowaction(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetAppflowaction(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListAppflowaction()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateAppflowaction(*resource)

	// assert.NoError(t, err)

	err = client.RenameAppflowaction(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteAppflowaction(resource.Name + "-rename")

	assert.NoError(t, err)
}
