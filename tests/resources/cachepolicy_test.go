package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCachepolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := Setup_cachepolicy(t, client)

	if resource == nil {
		return
	}

	key := resource.ToKey()

	log.Print("--ADD--")
	err := client.AddCachepolicy(*resource)

	assert.NoError(t, err)

	log.Print("--LIST--")
	list, err := client.ListCachepolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	log.Print("--GET--")
	res, err := client.GetCachepolicy(key)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--UPDATE--")
	err = client.UpdateCachepolicy(resource.ToUpdate())

	assert.NoError(t, err)

	log.Print("--COUNT--")
	count, err := client.CountCachepolicy()

	assert.NoError(t, err)
	assert.NotZero(t, count)

	log.Print("--EXISTS--")
	exists, err := client.ExistsCachepolicy(key)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--DELETE--")
	err = client.DeleteCachepolicy(key)

	assert.NoError(t, err)

	if tearDown != nil {
		tearDown()
	}
}
