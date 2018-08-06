package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFeoaction(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_feoaction(t, client)

	if resource == nil {
		return
	}

	err := client.AddFeoaction(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsFeoaction(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsFeoaction(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetFeoaction(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListFeoaction()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateFeoaction(*resource)

	// assert.NoError(t, err)

	err = client.RenameFeoaction(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteFeoaction(resource.Name + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
