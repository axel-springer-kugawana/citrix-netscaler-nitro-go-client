package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCspolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_cspolicy(t, client)

	if resource == nil {
		return
	}

	log.Print("--ADD--")
	err := client.AddCspolicy(*resource)

	assert.NoError(t, err)

	log.Print("--EXISTS--")
	exists, err := client.ExistsCspolicy(resource.Policyname + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	log.Print("--EXISTS--")
	exists, err = client.ExistsCspolicy(resource.Policyname)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--GET--")
	res, err := client.GetCspolicy(resource.Policyname)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--LIST--")
	list, err := client.ListCspolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateCspolicy(*resource)

	// assert.NoError(t, err)

	log.Print("--RENAME--")
	err = client.RenameCspolicy(resource.Policyname, resource.Policyname+"-rename")

	assert.NoError(t, err)

	log.Print("--DELETE--")
	err = client.DeleteCspolicy(resource.Policyname + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
