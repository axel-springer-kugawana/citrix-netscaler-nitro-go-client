package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVideooptimizationpolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_videooptimizationpolicy()

	if resource == nil {
		return
	}

	err := client.AddVideooptimizationpolicy(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsVideooptimizationpolicy(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err := client.ExistsVideooptimizationpolicy(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetVideooptimizationpolicy(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListVideooptimizationpolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateVideooptimizationpolicy(*resource)

	// assert.NoError(t, err)

	err = client.RenameVideooptimizationpolicy(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteVideooptimizationpolicy(resource.Name + "-rename")

	assert.NoError(t, err)
}
