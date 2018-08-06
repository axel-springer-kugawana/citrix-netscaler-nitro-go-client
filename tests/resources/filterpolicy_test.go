package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestFilterpolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_filterpolicy(t, client)

	if resource == nil {
		return
	}

	log.Print("--ADD--")
	err := client.AddFilterpolicy(*resource)

	assert.NoError(t, err)

	log.Print("--EXISTS--")
	exists, err := client.ExistsFilterpolicy(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	log.Print("--EXISTS--")
	exists, err = client.ExistsFilterpolicy(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--GET--")
	res, err := client.GetFilterpolicy(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--LIST--")
	list, err := client.ListFilterpolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateFilterpolicy(*resource)

	// assert.NoError(t, err)

	log.Print("--RENAME--")
	err = client.RenameFilterpolicy(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	log.Print("--DELETE--")
	err = client.DeleteFilterpolicy(resource.Name + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
