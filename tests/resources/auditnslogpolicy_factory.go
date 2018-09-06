package resources

import (
	"testing"

	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
)

func Setup_auditnslogpolicy(t *testing.T, client *nitro.NitroClient) (*nitro.Auditnslogpolicy, func()) {

	auditnslogaction, auditnslogactionTearDown := Setup_auditnslogaction(t, client)

	client.AddAuditnslogaction(*auditnslogaction)

	resource := nitro.Auditnslogpolicy{
		Name:   "auditnslogpolicy",
		Action: auditnslogaction.Name,
		Rule:   "ns_true",
	}
	return &resource, func() {
		client.DeleteAuditnslogaction(auditnslogaction.ToKey())

		if auditnslogactionTearDown != nil {
			auditnslogactionTearDown()
		}
	}
}
