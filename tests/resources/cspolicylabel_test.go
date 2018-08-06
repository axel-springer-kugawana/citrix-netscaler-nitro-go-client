package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCspolicylabel(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_cspolicylabel()

	if resource == nil {
		return
	}

	err := client.AddCspolicylabel(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsCspolicylabel(resource.Labelname + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err := client.ExistsCspolicylabel(resource.Labelname)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetCspolicylabel(resource.Labelname)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListCspolicylabel()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateCspolicylabel(*resource)

	// assert.NoError(t, err)

	err = client.RenameCspolicylabel(resource.Labelname, resource.Labelname+"-rename")

	assert.NoError(t, err)

	err = client.DeleteCspolicylabel(resource.Labelname + "-rename")

	assert.NoError(t, err)
}
