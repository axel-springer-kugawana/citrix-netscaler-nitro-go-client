package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLbmetrictable(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_lbmetrictable()

	if resource == nil {
		return
	}

	err := client.AddLbmetrictable(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsLbmetrictable(resource.Metrictable + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err := client.ExistsLbmetrictable(resource.Metrictable)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetLbmetrictable(resource.Metrictable)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListLbmetrictable()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateLbmetrictable(*resource)

	// assert.NoError(t, err)

	err = client.RenameLbmetrictable(resource.Metrictable, resource.Metrictable+"-rename")

	assert.NoError(t, err)

	err = client.DeleteLbmetrictable(resource.Metrictable + "-rename")

	assert.NoError(t, err)
}
