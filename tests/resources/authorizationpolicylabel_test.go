package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthorizationpolicylabel(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_authorizationpolicylabel()

	if resource == nil {
		return
	}

	err := client.AddAuthorizationpolicylabel(*resource)

	assert.NoError(t, err)

	res, err := client.GetAuthorizationpolicylabel(resource.Labelname)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListAuthorizationpolicylabel()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateAuthorizationpolicylabel(*resource)

	// assert.NoError(t, err)

	err = client.RenameAuthorizationpolicylabel(resource.Labelname, resource.Labelname+"-rename")

	assert.NoError(t, err)

	err = client.DeleteAuthorizationpolicylabel(resource.Labelname + "-rename")

	assert.NoError(t, err)
}