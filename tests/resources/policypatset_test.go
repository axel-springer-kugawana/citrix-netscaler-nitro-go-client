package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPolicypatset(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_policypatset()

	if resource == nil {
		return
	}

	err := client.AddPolicypatset(*resource)

	assert.NoError(t, err)

	res, err := client.GetPolicypatset(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListPolicypatset()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdatePolicypatset(*resource)

	// assert.NoError(t, err)

	err = client.RenamePolicypatset(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeletePolicypatset(resource.Name + "-rename")

	assert.NoError(t, err)
}
