package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppflowcollector(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_appflowcollector()

	if resource == nil {
		return
	}

	err := client.AddAppflowcollector(*resource)

	assert.NoError(t, err)

	res, err := client.GetAppflowcollector(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListAppflowcollector()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateAppflowcollector(*resource)

	// assert.NoError(t, err)

	err = client.RenameAppflowcollector(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteAppflowcollector(resource.Name + "-rename")

	assert.NoError(t, err)
}
