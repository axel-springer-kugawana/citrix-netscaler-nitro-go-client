package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPolicydataset(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_policydataset()

	if resource == nil {
		return
	}

	err := client.AddPolicydataset(*resource)

	assert.NoError(t, err)

	res, err := client.GetPolicydataset(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListPolicydataset()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdatePolicydataset(*resource)

	// assert.NoError(t, err)

	err = client.RenamePolicydataset(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeletePolicydataset(resource.Name + "-rename")

	assert.NoError(t, err)
}
