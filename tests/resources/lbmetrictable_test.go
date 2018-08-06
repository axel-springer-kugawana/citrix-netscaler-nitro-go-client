package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestLbmetrictable(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_lbmetrictable(t, client)

	if resource == nil {
		return
	}

	log.Print("--ADD--")
	err := client.AddLbmetrictable(*resource)

	assert.NoError(t, err)

	log.Print("--EXISTS--")
	exists, err := client.ExistsLbmetrictable(resource.Metrictable + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	log.Print("--EXISTS--")
	exists, err = client.ExistsLbmetrictable(resource.Metrictable)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--GET--")
	res, err := client.GetLbmetrictable(resource.Metrictable)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--LIST--")
	list, err := client.ListLbmetrictable()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateLbmetrictable(*resource)

	// assert.NoError(t, err)

	log.Print("--RENAME--")
	err = client.RenameLbmetrictable(resource.Metrictable, resource.Metrictable+"-rename")

	assert.NoError(t, err)

	log.Print("--DELETE--")
	err = client.DeleteLbmetrictable(resource.Metrictable + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
