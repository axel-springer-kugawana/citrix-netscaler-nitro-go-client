package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestAuthorizationpolicylabel(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_authorizationpolicylabel(t, client)

	if resource == nil {
		return
	}

	log.Print("--ADD--")
	err := client.AddAuthorizationpolicylabel(*resource)

	assert.NoError(t, err)

	log.Print("--EXISTS--")
	exists, err := client.ExistsAuthorizationpolicylabel(resource.Labelname + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	log.Print("--EXISTS--")
	exists, err = client.ExistsAuthorizationpolicylabel(resource.Labelname)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--GET--")
	res, err := client.GetAuthorizationpolicylabel(resource.Labelname)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--LIST--")
	list, err := client.ListAuthorizationpolicylabel()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateAuthorizationpolicylabel(*resource)

	// assert.NoError(t, err)

	log.Print("--RENAME--")
	err = client.RenameAuthorizationpolicylabel(resource.Labelname, resource.Labelname+"-rename")

	assert.NoError(t, err)

	log.Print("--DELETE--")
	err = client.DeleteAuthorizationpolicylabel(resource.Labelname + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
