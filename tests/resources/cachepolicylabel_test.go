package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCachepolicylabel(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_cachepolicylabel()

	if resource == nil {
		return
	}

	err := client.AddCachepolicylabel(*resource)

	assert.NoError(t, err)

	res, err := client.GetCachepolicylabel(resource.Labelname)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListCachepolicylabel()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateCachepolicylabel(*resource)

	// assert.NoError(t, err)

	err = client.RenameCachepolicylabel(resource.Labelname, resource.Labelname+"-rename")

	assert.NoError(t, err)

	err = client.DeleteCachepolicylabel(resource.Labelname + "-rename")

	assert.NoError(t, err)
}