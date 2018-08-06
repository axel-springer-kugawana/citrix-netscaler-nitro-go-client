package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestDnsaction64(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_dnsaction64(t, client)

	if resource == nil {
		return
	}

	log.Print("--ADD--")
	err := client.AddDnsaction64(*resource)

	assert.NoError(t, err)

	log.Print("--EXISTS--")
	exists, err := client.ExistsDnsaction64(resource.Actionname + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	log.Print("--EXISTS--")
	exists, err = client.ExistsDnsaction64(resource.Actionname)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--GET--")
	res, err := client.GetDnsaction64(resource.Actionname)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--LIST--")
	list, err := client.ListDnsaction64()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateDnsaction64(*resource)

	// assert.NoError(t, err)

	log.Print("--RENAME--")
	err = client.RenameDnsaction64(resource.Actionname, resource.Actionname+"-rename")

	assert.NoError(t, err)

	log.Print("--DELETE--")
	err = client.DeleteDnsaction64(resource.Actionname + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
