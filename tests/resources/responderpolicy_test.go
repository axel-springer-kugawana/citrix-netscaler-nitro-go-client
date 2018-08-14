package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestResponderpolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := Setup_responderpolicy(t, client)

	if resource == nil {
		return
	}

	key := resource.ToKey()

	log.Print("--ADD--")
	err := client.AddResponderpolicy(*resource)

	assert.NoError(t, err)

	log.Print("--LIST--")
	list, err := client.ListResponderpolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	log.Print("--GET--")
	res, err := client.GetResponderpolicy(key)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--UPDATE--")
	err = client.UpdateResponderpolicy(resource.ToUpdate())

	assert.NoError(t, err)

	log.Print("--COUNT--")
	count, err := client.CountResponderpolicy()

	assert.NoError(t, err)
	assert.NotZero(t, count)

	log.Print("--EXISTS--")
	exists, err := client.ExistsResponderpolicy(key)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--DELETE--")
	err = client.DeleteResponderpolicy(key)

	assert.NoError(t, err)

	if tearDown != nil {
		tearDown()
	}
}
