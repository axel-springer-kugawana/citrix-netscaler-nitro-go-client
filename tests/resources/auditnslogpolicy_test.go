package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuditnslogpolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_auditnslogpolicy()

	if resource == nil {
		return
	}

	err := client.AddAuditnslogpolicy(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsAuditnslogpolicy(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err := client.ExistsAuditnslogpolicy(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetAuditnslogpolicy(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListAuditnslogpolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateAuditnslogpolicy(*resource)

	// assert.NoError(t, err)

	err = client.RenameAuditnslogpolicy(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteAuditnslogpolicy(resource.Name + "-rename")

	assert.NoError(t, err)
}
