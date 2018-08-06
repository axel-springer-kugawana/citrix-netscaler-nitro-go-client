package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPolicystringmap(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_policystringmap(t, client)

	if resource == nil {
		return
	}

	err := client.AddPolicystringmap(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsPolicystringmap(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsPolicystringmap(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetPolicystringmap(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListPolicystringmap()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdatePolicystringmap(*resource)

	// assert.NoError(t, err)

	err = client.RenamePolicystringmap(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeletePolicystringmap(resource.Name + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
