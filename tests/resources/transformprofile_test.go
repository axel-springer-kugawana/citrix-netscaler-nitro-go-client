package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestTransformprofile(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_transformprofile(t, client)

	if resource == nil {
		return
	}

	key := resource.ToKey()

	log.Print("--ADD--")
	err := client.AddTransformprofile(*resource)

	assert.NoError(t, err)

	log.Print("--LIST--")
	list, err := client.ListTransformprofile()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	log.Print("--GET--")
	res, err := client.GetTransformprofile(key)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--COUNT--")
	count, err := client.CountTransformprofile()

	assert.NoError(t, err)
	assert.NotZero(t, count)

	log.Print("--EXISTS--")
	exists, err := client.ExistsTransformprofile(key)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--DELETE--")
	err = client.DeleteTransformprofile(key)

	assert.NoError(t, err)

	if tearDown != nil {
		tearDown()
	}
}
