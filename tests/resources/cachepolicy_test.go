package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCachepolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_cachepolicy()

	if resource == nil {
		return
	}

	err := client.AddCachepolicy(*resource)

	assert.NoError(t, err)

	res, err := client.GetCachepolicy(resource.Policyname)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListCachepolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateCachepolicy(*resource)

	// assert.NoError(t, err)

	err = client.RenameCachepolicy(resource.Policyname, resource.Policyname+"-rename")

	assert.NoError(t, err)

	err = client.DeleteCachepolicy(resource.Policyname + "-rename")

	assert.NoError(t, err)
}