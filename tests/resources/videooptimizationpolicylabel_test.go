package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestVideooptimizationpolicylabel(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_videooptimizationpolicylabel(t, client)

	if resource == nil {
		return
	}

	log.Print("--ADD--")
	err := client.AddVideooptimizationpolicylabel(*resource)

	assert.NoError(t, err)

	log.Print("--EXISTS--")
	exists, err := client.ExistsVideooptimizationpolicylabel(resource.Labelname + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	log.Print("--EXISTS--")
	exists, err = client.ExistsVideooptimizationpolicylabel(resource.Labelname)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--GET--")
	res, err := client.GetVideooptimizationpolicylabel(resource.Labelname)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--LIST--")
	list, err := client.ListVideooptimizationpolicylabel()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateVideooptimizationpolicylabel(*resource)

	// assert.NoError(t, err)

	log.Print("--RENAME--")
	err = client.RenameVideooptimizationpolicylabel(resource.Labelname, resource.Labelname+"-rename")

	assert.NoError(t, err)

	log.Print("--DELETE--")
	err = client.DeleteVideooptimizationpolicylabel(resource.Labelname + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
