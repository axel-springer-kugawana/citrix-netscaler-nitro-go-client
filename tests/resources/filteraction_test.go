package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilteraction(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := setup_filteraction(t)

	if resource == nil {
		return
	}

	err := client.AddFilteraction(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsFilteraction(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsFilteraction(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetFilteraction(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListFilteraction()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateFilteraction(*resource)

	// assert.NoError(t, err)

	err = client.RenameFilteraction(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteFilteraction(resource.Name + "-rename")

	assert.NoError(t, err)
}
