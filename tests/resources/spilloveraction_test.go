package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpilloveraction(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_spilloveraction(t, client)

	if resource == nil {
		return
	}

	err := client.AddSpilloveraction(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsSpilloveraction(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsSpilloveraction(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetSpilloveraction(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListSpilloveraction()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateSpilloveraction(*resource)

	// assert.NoError(t, err)

	err = client.RenameSpilloveraction(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteSpilloveraction(resource.Name + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
