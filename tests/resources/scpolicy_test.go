package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestScpolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := Setup_scpolicy(t, client)

	if resource == nil {
		return
	}

	key := resource.ToKey()

	log.Print("--ADD--")
	err := client.AddScpolicy(*resource)

	assert.NoError(t, err)

	log.Print("--LIST--")
	list, err := client.ListScpolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	log.Print("--GET--")
	res, err := client.GetScpolicy(key)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--COUNT--")
	count, err := client.CountScpolicy()

	assert.NoError(t, err)
	assert.NotZero(t, count)

	log.Print("--EXISTS--")
	exists, err := client.ExistsScpolicy(key)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--DELETE--")
	err = client.DeleteScpolicy(key)

	assert.NoError(t, err)

	if tearDown != nil {
		tearDown()
	}
}
