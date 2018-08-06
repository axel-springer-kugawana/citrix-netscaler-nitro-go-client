package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestLbprofile(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_lbprofile(t, client)

	if resource == nil {
		return
	}

	log.Print("--ADD--")
	err := client.AddLbprofile(*resource)

	assert.NoError(t, err)

	log.Print("--EXISTS--")
	exists, err := client.ExistsLbprofile(resource.Lbprofilename + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	log.Print("--EXISTS--")
	exists, err = client.ExistsLbprofile(resource.Lbprofilename)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--GET--")
	res, err := client.GetLbprofile(resource.Lbprofilename)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--LIST--")
	list, err := client.ListLbprofile()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateLbprofile(*resource)

	// assert.NoError(t, err)

	log.Print("--RENAME--")
	err = client.RenameLbprofile(resource.Lbprofilename, resource.Lbprofilename+"-rename")

	assert.NoError(t, err)

	log.Print("--DELETE--")
	err = client.DeleteLbprofile(resource.Lbprofilename + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
