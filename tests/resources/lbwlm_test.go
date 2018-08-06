package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestLbwlm(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_lbwlm(t, client)

	if resource == nil {
		return
	}

	log.Print("--ADD--")
	err := client.AddLbwlm(*resource)

	assert.NoError(t, err)

	log.Print("--EXISTS--")
	exists, err := client.ExistsLbwlm(resource.Wlmname + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	log.Print("--EXISTS--")
	exists, err = client.ExistsLbwlm(resource.Wlmname)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--GET--")
	res, err := client.GetLbwlm(resource.Wlmname)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--LIST--")
	list, err := client.ListLbwlm()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateLbwlm(*resource)

	// assert.NoError(t, err)

	log.Print("--RENAME--")
	err = client.RenameLbwlm(resource.Wlmname, resource.Wlmname+"-rename")

	assert.NoError(t, err)

	log.Print("--DELETE--")
	err = client.DeleteLbwlm(resource.Wlmname + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
