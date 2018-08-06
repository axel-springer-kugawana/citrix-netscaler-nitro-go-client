package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVideooptimizationaction(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_videooptimizationaction(t, client)

	if resource == nil {
		return
	}

	err := client.AddVideooptimizationaction(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsVideooptimizationaction(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsVideooptimizationaction(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetVideooptimizationaction(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListVideooptimizationaction()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateVideooptimizationaction(*resource)

	// assert.NoError(t, err)

	err = client.RenameVideooptimizationaction(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteVideooptimizationaction(resource.Name + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
