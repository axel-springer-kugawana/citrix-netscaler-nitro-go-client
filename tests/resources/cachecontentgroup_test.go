package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCachecontentgroup(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := setup_cachecontentgroup(t)

	if resource == nil {
		return
	}

	err := client.AddCachecontentgroup(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsCachecontentgroup(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsCachecontentgroup(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetCachecontentgroup(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListCachecontentgroup()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateCachecontentgroup(*resource)

	// assert.NoError(t, err)

	err = client.RenameCachecontentgroup(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteCachecontentgroup(resource.Name + "-rename")

	assert.NoError(t, err)
}
