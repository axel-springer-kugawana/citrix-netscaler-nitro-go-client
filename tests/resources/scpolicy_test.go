package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScpolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_scpolicy()

	if resource == nil {
		return
	}

	err := client.AddScpolicy(*resource)

	assert.NoError(t, err)

	res, err := client.GetScpolicy(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListScpolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateScpolicy(*resource)

	// assert.NoError(t, err)

	err = client.RenameScpolicy(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteScpolicy(resource.Name + "-rename")

	assert.NoError(t, err)
}
