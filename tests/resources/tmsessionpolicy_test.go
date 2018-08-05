package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTmsessionpolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_tmsessionpolicy()

	if resource == nil {
		return
	}

	err := client.AddTmsessionpolicy(*resource)

	assert.NoError(t, err)

	res, err := client.GetTmsessionpolicy(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListTmsessionpolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateTmsessionpolicy(*resource)

	// assert.NoError(t, err)

	err = client.RenameTmsessionpolicy(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteTmsessionpolicy(resource.Name + "-rename")

	assert.NoError(t, err)
}
