package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTmtrafficpolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_tmtrafficpolicy()

	if resource == nil {
		return
	}

	err := client.AddTmtrafficpolicy(*resource)

	assert.NoError(t, err)

	res, err := client.GetTmtrafficpolicy(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListTmtrafficpolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateTmtrafficpolicy(*resource)

	// assert.NoError(t, err)

	err = client.RenameTmtrafficpolicy(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteTmtrafficpolicy(resource.Name + "-rename")

	assert.NoError(t, err)
}