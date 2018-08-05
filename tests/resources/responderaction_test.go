package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResponderaction(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_responderaction()

	if resource == nil {
		return
	}

	err := client.AddResponderaction(*resource)

	assert.NoError(t, err)

	res, err := client.GetResponderaction(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListResponderaction()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateResponderaction(*resource)

	// assert.NoError(t, err)

	err = client.RenameResponderaction(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteResponderaction(resource.Name + "-rename")

	assert.NoError(t, err)
}
