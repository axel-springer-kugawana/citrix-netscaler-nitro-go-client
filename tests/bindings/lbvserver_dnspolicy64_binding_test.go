package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestLbvserverDnspolicy64Binding(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_lbvserver_dnspolicy64_binding(t, client)

	if resource == nil {
		return
	}

	key := resource.ToKey()

	log.Print("--ADD--")
	err := client.AddLbvserverDnspolicy64Binding(*resource)

	assert.NoError(t, err)

	log.Print("--LIST--")
	list, err := client.ListLbvserverDnspolicy64Binding()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	log.Print("--GET--")
	res, err := client.GetLbvserverDnspolicy64Binding(key)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--COUNT--")
	count, err := client.CountLbvserverDnspolicy64Binding()

	assert.NoError(t, err)
	assert.NotZero(t, count)

	log.Print("--EXISTS--")
	exists, err := client.ExistsLbvserverDnspolicy64Binding(key)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--DELETE--")
	err = client.DeleteLbvserverDnspolicy64Binding(key)

	assert.NoError(t, err)

	if tearDown != nil {
		tearDown()
	}
}
