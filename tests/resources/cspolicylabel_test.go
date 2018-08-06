package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCspolicylabel(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_cspolicylabel(t, client)

	if resource == nil {
		return
	}

	log.Print("--ADD--")
	err := client.AddCspolicylabel(*resource)

	assert.NoError(t, err)

	log.Print("--EXISTS--")
	exists, err := client.ExistsCspolicylabel(resource.Labelname + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	log.Print("--EXISTS--")
	exists, err = client.ExistsCspolicylabel(resource.Labelname)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--GET--")
	res, err := client.GetCspolicylabel(resource.Labelname)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--LIST--")
	list, err := client.ListCspolicylabel()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateCspolicylabel(*resource)

	// assert.NoError(t, err)

	log.Print("--RENAME--")
	err = client.RenameCspolicylabel(resource.Labelname, resource.Labelname+"-rename")

	assert.NoError(t, err)

	log.Print("--DELETE--")
	err = client.DeleteCspolicylabel(resource.Labelname + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
