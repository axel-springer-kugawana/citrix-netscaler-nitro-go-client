package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthorizationpolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := setup_authorizationpolicy(t)

	if resource == nil {
		return
	}

	err := client.AddAuthorizationpolicy(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsAuthorizationpolicy(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsAuthorizationpolicy(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetAuthorizationpolicy(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListAuthorizationpolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateAuthorizationpolicy(*resource)

	// assert.NoError(t, err)

	err = client.RenameAuthorizationpolicy(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteAuthorizationpolicy(resource.Name + "-rename")

	assert.NoError(t, err)
}
