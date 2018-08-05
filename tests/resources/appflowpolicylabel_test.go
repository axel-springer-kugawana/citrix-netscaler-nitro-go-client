package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppflowpolicylabel(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_appflowpolicylabel()

	if resource == nil {
		return
	}

	err := client.AddAppflowpolicylabel(*resource)

	assert.NoError(t, err)

	res, err := client.GetAppflowpolicylabel(resource.Labelname)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListAppflowpolicylabel()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateAppflowpolicylabel(*resource)

	// assert.NoError(t, err)

	err = client.RenameAppflowpolicylabel(resource.Labelname, resource.Labelname+"-rename")

	assert.NoError(t, err)

	err = client.DeleteAppflowpolicylabel(resource.Labelname + "-rename")

	assert.NoError(t, err)
}
