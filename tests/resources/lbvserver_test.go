package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestLbvserver(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_lbvserver(t, client)

	if resource == nil {
		return
	}

	log.Print("--ADD--")
	err := client.AddLbvserver(*resource)

	assert.NoError(t, err)

	log.Print("--EXISTS--")
	exists, err := client.ExistsLbvserver(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	log.Print("--EXISTS--")
	exists, err = client.ExistsLbvserver(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--GET--")
	res, err := client.GetLbvserver(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--LIST--")
	list, err := client.ListLbvserver()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateLbvserver(*resource)

	// assert.NoError(t, err)

	log.Print("--ENABLE--")
	err = client.EnableLbvserver(resource.Name)

	assert.NoError(t, err)

	log.Print("--DISABLE--")
	err = client.DisableLbvserver(resource.Name)

	assert.NoError(t, err)
	log.Print("--RENAME--")
	err = client.RenameLbvserver(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	log.Print("--DELETE--")
	err = client.DeleteLbvserver(resource.Name + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
