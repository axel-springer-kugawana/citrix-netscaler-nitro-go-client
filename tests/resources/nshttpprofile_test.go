package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestNshttpprofile(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_nshttpprofile(t, client)

	if resource == nil {
		return
	}

	log.Print("--ADD--")
	err := client.AddNshttpprofile(*resource)

	assert.NoError(t, err)

	log.Print("--EXISTS--")
	exists, err := client.ExistsNshttpprofile(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	log.Print("--EXISTS--")
	exists, err = client.ExistsNshttpprofile(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--GET--")
	res, err := client.GetNshttpprofile(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--LIST--")
	list, err := client.ListNshttpprofile()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateNshttpprofile(*resource)

	// assert.NoError(t, err)

	log.Print("--RENAME--")
	err = client.RenameNshttpprofile(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	log.Print("--DELETE--")
	err = client.DeleteNshttpprofile(resource.Name + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
