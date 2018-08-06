package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServicegroup(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_servicegroup()

	if resource == nil {
		return
	}

	err := client.AddServicegroup(*resource)

	assert.NoError(t, err)

	res, err := client.GetServicegroup(resource.Servicegroupname)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListServicegroup()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateServicegroup(*resource)

	// assert.NoError(t, err)

	err = client.RenameServicegroup(resource.Servicegroupname, resource.Servicegroupname+"-rename")

	assert.NoError(t, err)

	err = client.DeleteServicegroup(resource.Servicegroupname + "-rename")

	assert.NoError(t, err)
}