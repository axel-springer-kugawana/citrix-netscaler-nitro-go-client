package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestPqpolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_pqpolicy(t, client)

	if resource == nil {
		return
	}

	log.Print("--ADD--")
	err := client.AddPqpolicy(*resource)

	assert.NoError(t, err)

	log.Print("--EXISTS--")
	exists, err := client.ExistsPqpolicy(resource.Policyname + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	log.Print("--EXISTS--")
	exists, err = client.ExistsPqpolicy(resource.Policyname)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--GET--")
	res, err := client.GetPqpolicy(resource.Policyname)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--LIST--")
	list, err := client.ListPqpolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdatePqpolicy(*resource)

	// assert.NoError(t, err)

	log.Print("--RENAME--")
	err = client.RenamePqpolicy(resource.Policyname, resource.Policyname+"-rename")

	assert.NoError(t, err)

	log.Print("--DELETE--")
	err = client.DeletePqpolicy(resource.Policyname + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
