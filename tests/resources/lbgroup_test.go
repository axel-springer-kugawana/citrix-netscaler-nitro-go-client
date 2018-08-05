package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLbgroup(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_lbgroup()

	if resource == nil {
		return
	}

	err := client.AddLbgroup(*resource)

	assert.NoError(t, err)

	res, err := client.GetLbgroup(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListLbgroup()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateLbgroup(*resource)

	// assert.NoError(t, err)

	err = client.RenameLbgroup(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteLbgroup(resource.Name + "-rename")

	assert.NoError(t, err)
}
