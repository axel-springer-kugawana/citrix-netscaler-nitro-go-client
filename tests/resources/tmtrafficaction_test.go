package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestTmtrafficaction(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := Setup_tmtrafficaction(t, client)

	if resource == nil {
		return
	}

	key := resource.ToKey()

	log.Print("--ADD--")
	err := client.AddTmtrafficaction(*resource)

	assert.NoError(t, err)

	log.Print("--LIST--")
	list, err := client.ListTmtrafficaction()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	log.Print("--GET--")
	res, err := client.GetTmtrafficaction(key)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--UPDATE--")
	err = client.UpdateTmtrafficaction(resource.ToUpdate())

	assert.NoError(t, err)

	log.Print("--COUNT--")
	count, err := client.CountTmtrafficaction()

	assert.NoError(t, err)
	assert.NotZero(t, count)

	log.Print("--EXISTS--")
	exists, err := client.ExistsTmtrafficaction(key)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--DELETE--")
	err = client.DeleteTmtrafficaction(key)

	assert.NoError(t, err)

	if tearDown != nil {
		tearDown()
	}
}
