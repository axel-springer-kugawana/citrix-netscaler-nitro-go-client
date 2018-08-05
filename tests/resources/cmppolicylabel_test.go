package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCmppolicylabel(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_cmppolicylabel()

	if resource == nil {
		return
	}

	err := client.AddCmppolicylabel(*resource)

	assert.NoError(t, err)

	res, err := client.GetCmppolicylabel(resource.Labelname)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListCmppolicylabel()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateCmppolicylabel(*resource)

	// assert.NoError(t, err)

	err = client.RenameCmppolicylabel(resource.Labelname, resource.Labelname+"-rename")

	assert.NoError(t, err)

	err = client.DeleteCmppolicylabel(resource.Labelname + "-rename")

	assert.NoError(t, err)
}
