package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestAppqoepolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_appqoepolicy(t, client)

	if resource == nil {
		return
	}

	log.Print("--ADD--")
	err := client.AddAppqoepolicy(*resource)

	assert.NoError(t, err)

	log.Print("--EXISTS--")
	exists, err := client.ExistsAppqoepolicy(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	log.Print("--EXISTS--")
	exists, err = client.ExistsAppqoepolicy(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--GET--")
	res, err := client.GetAppqoepolicy(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--LIST--")
	list, err := client.ListAppqoepolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateAppqoepolicy(*resource)

	// assert.NoError(t, err)

	log.Print("--RENAME--")
	err = client.RenameAppqoepolicy(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	log.Print("--DELETE--")
	err = client.DeleteAppqoepolicy(resource.Name + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
