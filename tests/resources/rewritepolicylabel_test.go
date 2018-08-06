package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRewritepolicylabel(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_rewritepolicylabel(t, client)

	if resource == nil {
		return
	}

	err := client.AddRewritepolicylabel(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsRewritepolicylabel(resource.Labelname + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsRewritepolicylabel(resource.Labelname)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetRewritepolicylabel(resource.Labelname)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListRewritepolicylabel()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateRewritepolicylabel(*resource)

	// assert.NoError(t, err)

	err = client.RenameRewritepolicylabel(resource.Labelname, resource.Labelname+"-rename")

	assert.NoError(t, err)

	err = client.DeleteRewritepolicylabel(resource.Labelname + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
