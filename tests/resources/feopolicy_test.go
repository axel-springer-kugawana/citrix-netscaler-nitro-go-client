package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFeopolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_feopolicy(t, client)

	if resource == nil {
		return
	}

	err := client.AddFeopolicy(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsFeopolicy(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsFeopolicy(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetFeopolicy(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListFeopolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateFeopolicy(*resource)

	// assert.NoError(t, err)

	err = client.RenameFeopolicy(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteFeopolicy(resource.Name + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
