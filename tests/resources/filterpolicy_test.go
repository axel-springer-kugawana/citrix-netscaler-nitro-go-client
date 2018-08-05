package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterpolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_filterpolicy()

	if resource == nil {
		return
	}

	err := client.AddFilterpolicy(*resource)

	assert.NoError(t, err)

	res, err := client.GetFilterpolicy(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListFilterpolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateFilterpolicy(*resource)

	// assert.NoError(t, err)

	err = client.RenameFilterpolicy(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteFilterpolicy(resource.Name + "-rename")

	assert.NoError(t, err)
}
