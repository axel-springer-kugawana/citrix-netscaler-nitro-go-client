package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPqpolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_pqpolicy()

	if resource == nil {
		return
	}

	err := client.AddPqpolicy(*resource)

	assert.NoError(t, err)

	res, err := client.GetPqpolicy(resource.Policyname)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListPqpolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdatePqpolicy(*resource)

	// assert.NoError(t, err)

	err = client.RenamePqpolicy(resource.Policyname, resource.Policyname+"-rename")

	assert.NoError(t, err)

	err = client.DeletePqpolicy(resource.Policyname + "-rename")

	assert.NoError(t, err)
}
