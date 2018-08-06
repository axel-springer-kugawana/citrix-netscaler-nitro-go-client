package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppqoepolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_appqoepolicy()

	if resource == nil {
		return
	}

	err := client.AddAppqoepolicy(*resource)

	assert.NoError(t, err)

	res, err := client.GetAppqoepolicy(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListAppqoepolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateAppqoepolicy(*resource)

	// assert.NoError(t, err)

	err = client.RenameAppqoepolicy(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteAppqoepolicy(resource.Name + "-rename")

	assert.NoError(t, err)
}