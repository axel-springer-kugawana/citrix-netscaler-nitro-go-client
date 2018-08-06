package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServer(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := setup_server(t)

	if resource == nil {
		return
	}

	err := client.AddServer(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsServer(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsServer(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetServer(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListServer()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateServer(*resource)

	// assert.NoError(t, err)

	err = client.EnableServer(resource.Name)

	assert.NoError(t, err)

	err = client.DisableServer(resource.Name)

	assert.NoError(t, err)
	err = client.RenameServer(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteServer(resource.Name + "-rename")

	assert.NoError(t, err)
}
