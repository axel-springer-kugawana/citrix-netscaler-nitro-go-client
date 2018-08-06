package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCapolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_capolicy()

	if resource == nil {
		return
	}

	err := client.AddCapolicy(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsCapolicy(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err := client.ExistsCapolicy(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetCapolicy(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListCapolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateCapolicy(*resource)

	// assert.NoError(t, err)

	err = client.RenameCapolicy(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteCapolicy(resource.Name + "-rename")

	assert.NoError(t, err)
}
