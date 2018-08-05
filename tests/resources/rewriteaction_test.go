package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRewriteaction(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_rewriteaction()

	if resource == nil {
		return
	}

	err := client.AddRewriteaction(*resource)

	assert.NoError(t, err)

	res, err := client.GetRewriteaction(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListRewriteaction()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateRewriteaction(*resource)

	// assert.NoError(t, err)

	err = client.RenameRewriteaction(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteRewriteaction(resource.Name + "-rename")

	assert.NoError(t, err)
}
