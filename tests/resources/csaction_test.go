package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCsaction(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := setup_csaction(t)

	if resource == nil {
		return
	}

	err := client.AddCsaction(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsCsaction(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsCsaction(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetCsaction(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListCsaction()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateCsaction(*resource)

	// assert.NoError(t, err)

	err = client.RenameCsaction(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteCsaction(resource.Name + "-rename")

	assert.NoError(t, err)
}
