package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPolicyexpression(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_policyexpression()

	if resource == nil {
		return
	}

	err := client.AddPolicyexpression(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsPolicyexpression(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err := client.ExistsPolicyexpression(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetPolicyexpression(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListPolicyexpression()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdatePolicyexpression(*resource)

	// assert.NoError(t, err)

	err = client.RenamePolicyexpression(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeletePolicyexpression(resource.Name + "-rename")

	assert.NoError(t, err)
}
