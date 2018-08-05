package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVideooptimizationpolicylabel(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_videooptimizationpolicylabel()

	if resource == nil {
		return
	}

	err := client.AddVideooptimizationpolicylabel(*resource)

	assert.NoError(t, err)

	res, err := client.GetVideooptimizationpolicylabel(resource.Labelname)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListVideooptimizationpolicylabel()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateVideooptimizationpolicylabel(*resource)

	// assert.NoError(t, err)

	err = client.RenameVideooptimizationpolicylabel(resource.Labelname, resource.Labelname+"-rename")

	assert.NoError(t, err)

	err = client.DeleteVideooptimizationpolicylabel(resource.Labelname + "-rename")

	assert.NoError(t, err)
}
