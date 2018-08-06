package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestDnsprofile(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_dnsprofile(t, client)

	if resource == nil {
		return
	}

	log.Print("--ADD--")
	err := client.AddDnsprofile(*resource)

	assert.NoError(t, err)

	log.Print("--EXISTS--")
	exists, err := client.ExistsDnsprofile(resource.Dnsprofilename + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	log.Print("--EXISTS--")
	exists, err = client.ExistsDnsprofile(resource.Dnsprofilename)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--GET--")
	res, err := client.GetDnsprofile(resource.Dnsprofilename)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--LIST--")
	list, err := client.ListDnsprofile()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateDnsprofile(*resource)

	// assert.NoError(t, err)

	log.Print("--RENAME--")
	err = client.RenameDnsprofile(resource.Dnsprofilename, resource.Dnsprofilename+"-rename")

	assert.NoError(t, err)

	log.Print("--DELETE--")
	err = client.DeleteDnsprofile(resource.Dnsprofilename + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
