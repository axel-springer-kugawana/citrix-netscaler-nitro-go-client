package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVideooptimizationaction(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_videooptimizationaction()

	if resource == nil {
		return
	}

	err := client.AddVideooptimizationaction(*resource)

	assert.NoError(t, err)

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
}
