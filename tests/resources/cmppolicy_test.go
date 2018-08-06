package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCmppolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_cmppolicy()

	if resource == nil {
		return
	}

	err := client.AddCmppolicy(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsCmppolicy(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err := client.ExistsCmppolicy(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetCmppolicy(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListCmppolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateCmppolicy(*resource)

	// assert.NoError(t, err)

	err = client.RenameCmppolicy(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteCmppolicy(resource.Name + "-rename")

	assert.NoError(t, err)
}
