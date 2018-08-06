package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDospolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := setup_dospolicy(t)

	if resource == nil {
		return
	}

	err := client.AddDospolicy(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsDospolicy(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsDospolicy(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetDospolicy(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListDospolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateDospolicy(*resource)

	// assert.NoError(t, err)

	err = client.RenameDospolicy(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteDospolicy(resource.Name + "-rename")

	assert.NoError(t, err)
}
