package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpilloverpolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_spilloverpolicy()

	if resource == nil {
		return
	}

	err := client.AddSpilloverpolicy(*resource)

	assert.NoError(t, err)

	res, err := client.GetSpilloverpolicy(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListSpilloverpolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateSpilloverpolicy(*resource)

	// assert.NoError(t, err)

	err = client.RenameSpilloverpolicy(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteSpilloverpolicy(resource.Name + "-rename")

	assert.NoError(t, err)
}
