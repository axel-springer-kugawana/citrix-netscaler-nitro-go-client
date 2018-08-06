package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTmtrafficaction(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_tmtrafficaction()

	if resource == nil {
		return
	}

	err := client.AddTmtrafficaction(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsTmtrafficaction(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err := client.ExistsTmtrafficaction(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetTmtrafficaction(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListTmtrafficaction()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateTmtrafficaction(*resource)

	// assert.NoError(t, err)

	err = client.RenameTmtrafficaction(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteTmtrafficaction(resource.Name + "-rename")

	assert.NoError(t, err)
}
