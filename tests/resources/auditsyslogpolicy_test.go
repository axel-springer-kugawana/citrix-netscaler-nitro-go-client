package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuditsyslogpolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_auditsyslogpolicy()

	if resource == nil {
		return
	}

	err := client.AddAuditsyslogpolicy(*resource)

	assert.NoError(t, err)

	res, err := client.GetAuditsyslogpolicy(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListAuditsyslogpolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateAuditsyslogpolicy(*resource)

	// assert.NoError(t, err)

	err = client.RenameAuditsyslogpolicy(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteAuditsyslogpolicy(resource.Name + "-rename")

	assert.NoError(t, err)
}
