package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppfwpolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := setup_appfwpolicy(t)

	if resource == nil {
		return
	}

	err := client.AddAppfwpolicy(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsAppfwpolicy(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsAppfwpolicy(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetAppfwpolicy(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListAppfwpolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateAppfwpolicy(*resource)

	// assert.NoError(t, err)

	err = client.RenameAppfwpolicy(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteAppfwpolicy(resource.Name + "-rename")

	assert.NoError(t, err)
}
