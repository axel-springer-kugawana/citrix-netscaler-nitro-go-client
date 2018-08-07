package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestFilteraction(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_filteraction(t, client)

	if resource == nil {
		return
	}

	key := resource.ToKey()

	log.Print("--ADD--")
	err := client.AddFilteraction(*resource)

	assert.NoError(t, err)

	log.Print("--LIST--")
	list, err := client.ListFilteraction()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	log.Print("--GET--")
	res, err := client.GetFilteraction(key)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--COUNT--")
	count, err := client.CountFilteraction()

	assert.NoError(t, err)
	assert.NotZero(t, count)

	log.Print("--EXISTS--")
	exists, err := client.ExistsFilteraction(key)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--DELETE--")
	err = client.DeleteFilteraction(key)

	assert.NoError(t, err)

	if tearDown != nil {
		tearDown()
	}
}
