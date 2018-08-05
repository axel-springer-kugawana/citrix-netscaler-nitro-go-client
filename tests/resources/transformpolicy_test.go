package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransformpolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_transformpolicy()

	if resource == nil {
		return
	}

	err := client.AddTransformpolicy(*resource)

	assert.NoError(t, err)

	res, err := client.GetTransformpolicy(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListTransformpolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateTransformpolicy(*resource)

	// assert.NoError(t, err)

	err = client.RenameTransformpolicy(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteTransformpolicy(resource.Name + "-rename")

	assert.NoError(t, err)
}
