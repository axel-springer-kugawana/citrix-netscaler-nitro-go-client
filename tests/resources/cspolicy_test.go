package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCspolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_cspolicy()

	if resource == nil {
		return
	}

	err := client.AddCspolicy(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsCspolicy(resource.Policyname + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err := client.ExistsCspolicy(resource.Policyname)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetCspolicy(resource.Policyname)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListCspolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateCspolicy(*resource)

	// assert.NoError(t, err)

	err = client.RenameCspolicy(resource.Policyname, resource.Policyname+"-rename")

	assert.NoError(t, err)

	err = client.DeleteCspolicy(resource.Policyname + "-rename")

	assert.NoError(t, err)
}
