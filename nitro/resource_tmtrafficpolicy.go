package nitro

type Tmtrafficpolicy struct {
	Name   string `json:"name"`
	Action string `json:"action,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

type TmtrafficpolicyKey struct {
	Name string
}

type tmtrafficpolicy_update struct {
	Name   string `json:"name"`
	Rule   string `json:"rule,omitempty"`
	Action string `json:"action,omitempty"`
}

type tmtrafficpolicy_payload struct {
	tmtrafficpolicy interface{}
}

func tmtrafficpolicy_key_to_args(key TmtrafficpolicyKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteTmtrafficpolicy(key TmtrafficpolicyKey) error {
	return c.deleteResourceWithArgs("tmtrafficpolicy", key.Name, tmtrafficpolicy_key_to_args(key))
}

func (c *NitroClient) GetTmtrafficpolicy(key TmtrafficpolicyKey) (*Tmtrafficpolicy, error) {
	var results struct {
		Tmtrafficpolicy []Tmtrafficpolicy
	}

	if err := c.getResourceWithArgs("tmtrafficpolicy", key.Name, tmtrafficpolicy_key_to_args(key), &results); err != nil || len(results.Tmtrafficpolicy) != 1 {
		return nil, err
	}

	return &results.Tmtrafficpolicy[0], nil
}

func (c *NitroClient) ListTmtrafficpolicy() ([]Tmtrafficpolicy, error) {
	var results struct {
		Tmtrafficpolicy []Tmtrafficpolicy
	}

	if err := c.listResources("tmtrafficpolicy", &results); err != nil {
		return nil, err
	}

	return results.Tmtrafficpolicy, nil
}

func (c *NitroClient) AddTmtrafficpolicy(resource Tmtrafficpolicy) error {
	return c.addResource("tmtrafficpolicy", resource)
}

func (c *NitroClient) RenameTmtrafficpolicy(name string, newName string) error {
	return c.renameResource("tmtrafficpolicy", "name", name, newName)
}

func (c *NitroClient) UnsetTmtrafficpolicy(name string, fields ...string) error {
	return c.unsetResource("tmtrafficpolicy", "name", name, fields)
}

func (c *NitroClient) UpdateTmtrafficpolicy(resource Tmtrafficpolicy) error {
	update := tmtrafficpolicy_update{
		resource.Name,
		resource.Rule,
		resource.Action,
	}

	return c.Put("tmtrafficpolicy", update)
}
