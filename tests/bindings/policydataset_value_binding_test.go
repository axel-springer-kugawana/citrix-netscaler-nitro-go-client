package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestPolicydatasetValueBinding(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := Setup_policydataset_value_binding(t, client)

	if resource == nil {
		return
	}

	key := resource.ToKey()

	log.Print("--COUNT--")
	count0, err := client.CountPolicydatasetValueBinding()

	assert.NoError(t, err)

	log.Print("--ADD--")
	err = client.AddPolicydatasetValueBinding(*resource)

	assert.NoError(t, err)

	log.Print("--LIST--")
	list, err := client.ListPolicydatasetValueBinding()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	log.Print("--COUNT--")
	count, err := client.CountPolicydatasetValueBinding()

	assert.NoError(t, err)
	assert.NotZero(t, count)
	assert.Equal(t, count, count0+1)

	log.Print("--GET--")
	res, err := client.GetPolicydatasetValueBinding(key)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--EXISTS--")
	exists, err := client.ExistsPolicydatasetValueBinding(key)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--DELETE--")
	err = client.DeletePolicydatasetValueBinding(key)

	assert.NoError(t, err)

	if tearDown != nil {
		tearDown()
	}
}
