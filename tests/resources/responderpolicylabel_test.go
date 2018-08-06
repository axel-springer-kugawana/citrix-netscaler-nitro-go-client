package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResponderpolicylabel(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_responderpolicylabel(t, client)

	if resource == nil {
		return
	}

	err := client.AddResponderpolicylabel(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsResponderpolicylabel(resource.Labelname + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsResponderpolicylabel(resource.Labelname)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetResponderpolicylabel(resource.Labelname)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListResponderpolicylabel()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateResponderpolicylabel(*resource)

	// assert.NoError(t, err)

	err = client.RenameResponderpolicylabel(resource.Labelname, resource.Labelname+"-rename")

	assert.NoError(t, err)

	err = client.DeleteResponderpolicylabel(resource.Labelname + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
