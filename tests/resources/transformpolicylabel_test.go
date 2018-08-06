package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransformpolicylabel(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_transformpolicylabel(t, client)

	if resource == nil {
		return
	}

	err := client.AddTransformpolicylabel(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsTransformpolicylabel(resource.Labelname + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsTransformpolicylabel(resource.Labelname)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetTransformpolicylabel(resource.Labelname)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListTransformpolicylabel()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateTransformpolicylabel(*resource)

	// assert.NoError(t, err)

	err = client.RenameTransformpolicylabel(resource.Labelname, resource.Labelname+"-rename")

	assert.NoError(t, err)

	err = client.DeleteTransformpolicylabel(resource.Labelname + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
