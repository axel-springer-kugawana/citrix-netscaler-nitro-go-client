package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLbvserver(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_lbvserver()

	if resource == nil {
		return
	}

	err := client.AddLbvserver(*resource)

	assert.NoError(t, err)

	res, err := client.GetLbvserver(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListLbvserver()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateLbvserver(*resource)

	// assert.NoError(t, err)

	err = client.EnableLbvserver(resource.Name)

	assert.NoError(t, err)

	err = client.DisableLbvserver(resource.Name)

	assert.NoError(t, err)
	err = client.RenameLbvserver(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteLbvserver(resource.Name + "-rename")

	assert.NoError(t, err)
}
