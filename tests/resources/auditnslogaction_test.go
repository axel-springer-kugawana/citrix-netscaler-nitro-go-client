package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuditnslogaction(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource := create_auditnslogaction()

	if resource == nil {
		return
	}

	err := client.AddAuditnslogaction(*resource)

	assert.NoError(t, err)

	res, err := client.GetAuditnslogaction(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListAuditnslogaction()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateAuditnslogaction(*resource)

	// assert.NoError(t, err)

	err = client.RenameAuditnslogaction(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteAuditnslogaction(resource.Name + "-rename")

	assert.NoError(t, err)
}