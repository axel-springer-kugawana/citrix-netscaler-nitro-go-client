package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuditsyslogaction(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := setup_auditsyslogaction(t)

	if resource == nil {
		return
	}

	err := client.AddAuditsyslogaction(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsAuditsyslogaction(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsAuditsyslogaction(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetAuditsyslogaction(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListAuditsyslogaction()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateAuditsyslogaction(*resource)

	// assert.NoError(t, err)

	err = client.RenameAuditsyslogaction(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteAuditsyslogaction(resource.Name + "-rename")

	assert.NoError(t, err)
}
