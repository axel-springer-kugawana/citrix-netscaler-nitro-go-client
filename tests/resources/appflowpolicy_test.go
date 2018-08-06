package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppflowpolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_appflowpolicy(t, client)

	if resource == nil {
		return
	}

	err := client.AddAppflowpolicy(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsAppflowpolicy(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsAppflowpolicy(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetAppflowpolicy(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListAppflowpolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateAppflowpolicy(*resource)

	// assert.NoError(t, err)

	err = client.RenameAppflowpolicy(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteAppflowpolicy(resource.Name + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
