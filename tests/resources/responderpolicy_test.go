package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResponderpolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_responderpolicy()

	if resource == nil {
		return
	}

	err := client.AddResponderpolicy(*resource)

	assert.NoError(t, err)

	res, err := client.GetResponderpolicy(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListResponderpolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateResponderpolicy(*resource)

	// assert.NoError(t, err)

	err = client.RenameResponderpolicy(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteResponderpolicy(resource.Name + "-rename")

	assert.NoError(t, err)
}
