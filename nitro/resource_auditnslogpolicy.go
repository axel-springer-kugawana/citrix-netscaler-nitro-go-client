package nitro

type AuditnslogpolicyResource struct {
	Name   string `json:"name"`
	Action string `json:"action,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

type AuditnslogpolicyKey struct {
	Name string
}

type auditnslogpolicy_update struct {
	Name   string `json:"name"`
	Rule   string `json:"rule,omitempty"`
	Action string `json:"action,omitempty"`
}

type auditnslogpolicy_payload struct {
	auditnslogpolicy interface{}
}

func auditnslogpolicy_key_to_args(key AuditnslogpolicyKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteAuditnslogpolicy(key AuditnslogpolicyKey) error {
	return c.deleteResourceWithArgs("auditnslogpolicy", key.Name, auditnslogpolicy_key_to_args(key))
}

func (c *NitroClient) GetAuditnslogpolicy(key AuditnslogpolicyKey) (*AuditnslogpolicyResource, error) {
	var results struct {
		Auditnslogpolicy []AuditnslogpolicyResource
	}

	if err := c.getResourceWithArgs("auditnslogpolicy", key.Name, auditnslogpolicy_key_to_args(key), &results); err != nil || len(results.Auditnslogpolicy) != 1 {
		return nil, err
	}

	return &results.Auditnslogpolicy[0], nil
}

func (c *NitroClient) ListAuditnslogpolicy() ([]AuditnslogpolicyResource, error) {
	var results struct {
		Auditnslogpolicy []AuditnslogpolicyResource
	}

	if err := c.listResources("auditnslogpolicy", &results); err != nil {
		return nil, err
	}

	return results.Auditnslogpolicy, nil
}

func (c *NitroClient) AddAuditnslogpolicy(resource AuditnslogpolicyResource) error {
	return c.addResource("auditnslogpolicy", resource)
}

func (c *NitroClient) RenameAuditnslogpolicy(name string, newName string) error {
	return c.renameResource("auditnslogpolicy", "name", name, newName)
}

func (c *NitroClient) UnsetAuditnslogpolicy(name string, fields ...string) error {
	return c.unsetResource("auditnslogpolicy", "name", name, fields)
}

func (c *NitroClient) UpdateAuditnslogpolicy(resource AuditnslogpolicyResource) error {
	update := auditnslogpolicy_update{
		resource.Name,
		resource.Rule,
		resource.Action,
	}

	return c.Put("auditnslogpolicy", update)
}
