package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestVideooptimizationpolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_videooptimizationpolicy(t, client)

	if resource == nil {
		return
	}

	log.Print("--ADD--")
	err := client.AddVideooptimizationpolicy(*resource)

	assert.NoError(t, err)

	log.Print("--EXISTS--")
	exists, err := client.ExistsVideooptimizationpolicy(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	log.Print("--EXISTS--")
	exists, err = client.ExistsVideooptimizationpolicy(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--GET--")
	res, err := client.GetVideooptimizationpolicy(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--LIST--")
	list, err := client.ListVideooptimizationpolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateVideooptimizationpolicy(*resource)

	// assert.NoError(t, err)

	log.Print("--RENAME--")
	err = client.RenameVideooptimizationpolicy(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	log.Print("--DELETE--")
	err = client.DeleteVideooptimizationpolicy(resource.Name + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
