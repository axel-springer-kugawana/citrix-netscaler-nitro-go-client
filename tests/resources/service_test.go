package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_service()

	if resource == nil {
		return
	}

	err := client.AddService(*resource)

	assert.NoError(t, err)

	res, err := client.GetService(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListService()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateService(*resource)

	// assert.NoError(t, err)

	err = client.RenameService(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteService(resource.Name + "-rename")

	assert.NoError(t, err)
}
