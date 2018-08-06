package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTmsessionaction(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_tmsessionaction(t, client)

	if resource == nil {
		return
	}

	err := client.AddTmsessionaction(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsTmsessionaction(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsTmsessionaction(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetTmsessionaction(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListTmsessionaction()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateTmsessionaction(*resource)

	// assert.NoError(t, err)

	err = client.RenameTmsessionaction(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteTmsessionaction(resource.Name + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
