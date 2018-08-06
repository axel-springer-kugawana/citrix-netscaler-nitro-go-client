package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDbdbprofile(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_dbdbprofile()

	if resource == nil {
		return
	}

	err := client.AddDbdbprofile(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsDbdbprofile(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err := client.ExistsDbdbprofile(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetDbdbprofile(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListDbdbprofile()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateDbdbprofile(*resource)

	// assert.NoError(t, err)

	err = client.RenameDbdbprofile(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteDbdbprofile(resource.Name + "-rename")

	assert.NoError(t, err)
}
