package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDnspolicy64(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_dnspolicy64()

	if resource == nil {
		return
	}

	err := client.AddDnspolicy64(*resource)

	assert.NoError(t, err)

	res, err := client.GetDnspolicy64(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListDnspolicy64()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateDnspolicy64(*resource)

	// assert.NoError(t, err)

	err = client.RenameDnspolicy64(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteDnspolicy64(resource.Name + "-rename")

	assert.NoError(t, err)
}