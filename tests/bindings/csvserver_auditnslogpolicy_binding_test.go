package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCsvserverAuditnslogpolicyBinding(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := Setup_csvserver_auditnslogpolicy_binding(t, client)

	if resource == nil {
		return
	}

	key := resource.ToKey()

	log.Print("--COUNT--")
	count0, err := client.CountCsvserverAuditnslogpolicyBinding()

	assert.NoError(t, err)

	log.Print("--ADD--")
	err = client.AddCsvserverAuditnslogpolicyBinding(*resource)

	assert.NoError(t, err)

	log.Print("--LIST--")
	list, err := client.ListCsvserverAuditnslogpolicyBinding()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	log.Print("--COUNT--")
	count, err := client.CountCsvserverAuditnslogpolicyBinding()

	assert.NoError(t, err)
	assert.NotZero(t, count)
	assert.Equal(t, count, count0+1)

	log.Print("--GET--")
	res, err := client.GetCsvserverAuditnslogpolicyBinding(key)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--EXISTS--")
	exists, err := client.ExistsCsvserverAuditnslogpolicyBinding(key)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--DELETE--")
	err = client.DeleteCsvserverAuditnslogpolicyBinding(key)

	assert.NoError(t, err)

	if tearDown != nil {
		tearDown()
	}
}
