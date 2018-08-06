package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCsvserver(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_csvserver(t, client)

	if resource == nil {
		return
	}

	err := client.AddCsvserver(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsCsvserver(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsCsvserver(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetCsvserver(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListCsvserver()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateCsvserver(*resource)

	// assert.NoError(t, err)

	err = client.EnableCsvserver(resource.Name)

	assert.NoError(t, err)

	err = client.DisableCsvserver(resource.Name)

	assert.NoError(t, err)
	err = client.RenameCsvserver(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteCsvserver(resource.Name + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
