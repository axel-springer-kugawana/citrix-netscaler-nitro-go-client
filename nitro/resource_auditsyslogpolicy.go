package nitro

type Auditsyslogpolicy struct {
	Name   string `json:"name"`
	Action string `json:"action,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

type AuditsyslogpolicyKey struct {
	Name string
}

type auditsyslogpolicy_update struct {
	Name   string `json:"name"`
	Rule   string `json:"rule,omitempty"`
	Action string `json:"action,omitempty"`
}

type auditsyslogpolicy_payload struct {
	auditsyslogpolicy interface{}
}

func auditsyslogpolicy_key_to_args(key AuditsyslogpolicyKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteAuditsyslogpolicy(key AuditsyslogpolicyKey) error {
	return c.deleteResourceWithArgs("auditsyslogpolicy", key.Name, auditsyslogpolicy_key_to_args(key))
}

func (c *NitroClient) GetAuditsyslogpolicy(key AuditsyslogpolicyKey) (*Auditsyslogpolicy, error) {
	var results struct {
		Auditsyslogpolicy []Auditsyslogpolicy
	}

	if err := c.getResourceWithArgs("auditsyslogpolicy", key.Name, auditsyslogpolicy_key_to_args(key), &results); err != nil || len(results.Auditsyslogpolicy) != 1 {
		return nil, err
	}

	return &results.Auditsyslogpolicy[0], nil
}

func (c *NitroClient) ListAuditsyslogpolicy() ([]Auditsyslogpolicy, error) {
	var results struct {
		Auditsyslogpolicy []Auditsyslogpolicy
	}

	if err := c.listResources("auditsyslogpolicy", &results); err != nil {
		return nil, err
	}

	return results.Auditsyslogpolicy, nil
}

func (c *NitroClient) AddAuditsyslogpolicy(resource Auditsyslogpolicy) error {
	return c.addResource("auditsyslogpolicy", resource)
}

func (c *NitroClient) RenameAuditsyslogpolicy(name string, newName string) error {
	return c.renameResource("auditsyslogpolicy", "name", name, newName)
}

func (c *NitroClient) UnsetAuditsyslogpolicy(name string, fields ...string) error {
	return c.unsetResource("auditsyslogpolicy", "name", name, fields)
}

func (c *NitroClient) UpdateAuditsyslogpolicy(resource Auditsyslogpolicy) error {
	update := auditsyslogpolicy_update{
		resource.Name,
		resource.Rule,
		resource.Action,
	}

	return c.Put("auditsyslogpolicy", update)
}
