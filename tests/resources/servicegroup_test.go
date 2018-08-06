package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestServicegroup(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_servicegroup(t, client)

	if resource == nil {
		return
	}

	log.Print("--ADD--")
	err := client.AddServicegroup(*resource)

	assert.NoError(t, err)

	log.Print("--EXISTS--")
	exists, err := client.ExistsServicegroup(resource.Servicegroupname + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	log.Print("--EXISTS--")
	exists, err = client.ExistsServicegroup(resource.Servicegroupname)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--GET--")
	res, err := client.GetServicegroup(resource.Servicegroupname)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--LIST--")
	list, err := client.ListServicegroup()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateServicegroup(*resource)

	// assert.NoError(t, err)

	log.Print("--RENAME--")
	err = client.RenameServicegroup(resource.Servicegroupname, resource.Servicegroupname+"-rename")

	assert.NoError(t, err)

	log.Print("--DELETE--")
	err = client.DeleteServicegroup(resource.Servicegroupname + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
