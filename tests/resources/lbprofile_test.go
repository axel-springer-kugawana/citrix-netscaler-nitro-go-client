package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLbprofile(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := setup_lbprofile(t)

	if resource == nil {
		return
	}

	err := client.AddLbprofile(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsLbprofile(resource.Lbprofilename + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsLbprofile(resource.Lbprofilename)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetLbprofile(resource.Lbprofilename)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListLbprofile()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateLbprofile(*resource)

	// assert.NoError(t, err)

	err = client.RenameLbprofile(resource.Lbprofilename, resource.Lbprofilename+"-rename")

	assert.NoError(t, err)

	err = client.DeleteLbprofile(resource.Lbprofilename + "-rename")

	assert.NoError(t, err)
}
