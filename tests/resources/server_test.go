package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestServer(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_server(t, client)

	if resource == nil {
		return
	}

	log.Print("--ADD--")
	err := client.AddServer(*resource)

	assert.NoError(t, err)

	log.Print("--EXISTS--")
	exists, err := client.ExistsServer(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	log.Print("--EXISTS--")
	exists, err = client.ExistsServer(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--GET--")
	res, err := client.GetServer(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--LIST--")
	list, err := client.ListServer()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateServer(*resource)

	// assert.NoError(t, err)

	log.Print("--ENABLE--")
	err = client.EnableServer(resource.Name)

	assert.NoError(t, err)

	log.Print("--DISABLE--")
	err = client.DisableServer(resource.Name)

	assert.NoError(t, err)
	log.Print("--RENAME--")
	err = client.RenameServer(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	log.Print("--DELETE--")
	err = client.DeleteServer(resource.Name + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
