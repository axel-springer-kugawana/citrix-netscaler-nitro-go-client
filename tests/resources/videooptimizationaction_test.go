package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestVideooptimizationaction(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_videooptimizationaction(t, client)

	if resource == nil {
		return
	}

	log.Print("--ADD--")
	err := client.AddVideooptimizationaction(*resource)

	assert.NoError(t, err)

	log.Print("--EXISTS--")
	exists, err := client.ExistsVideooptimizationaction(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	log.Print("--EXISTS--")
	exists, err = client.ExistsVideooptimizationaction(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--GET--")
	res, err := client.GetVideooptimizationaction(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--LIST--")
	list, err := client.ListVideooptimizationaction()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateVideooptimizationaction(*resource)

	// assert.NoError(t, err)

	log.Print("--RENAME--")
	err = client.RenameVideooptimizationaction(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	log.Print("--DELETE--")
	err = client.DeleteVideooptimizationaction(resource.Name + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
