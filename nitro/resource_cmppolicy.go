package nitro

type Cmppolicy struct {
	Name      string `json:"name"`
	Resaction string `json:"resaction,omitempty"`
	Rule      string `json:"rule,omitempty"`
}

type CmppolicyKey struct {
	Name string
}

type cmppolicy_update struct {
	Name      string `json:"name"`
	Rule      string `json:"rule,omitempty"`
	Resaction string `json:"resaction,omitempty"`
}

type cmppolicy_payload struct {
	cmppolicy interface{}
}

func cmppolicy_key_to_args(key CmppolicyKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteCmppolicy(key CmppolicyKey) error {
	return c.deleteResourceWithArgs("cmppolicy", key.Name, cmppolicy_key_to_args(key))
}

func (c *NitroClient) GetCmppolicy(key CmppolicyKey) (*Cmppolicy, error) {
	var results struct {
		Cmppolicy []Cmppolicy
	}

	if err := c.getResourceWithArgs("cmppolicy", key.Name, cmppolicy_key_to_args(key), &results); err != nil || len(results.Cmppolicy) != 1 {
		return nil, err
	}

	return &results.Cmppolicy[0], nil
}

func (c *NitroClient) ListCmppolicy() ([]Cmppolicy, error) {
	var results struct {
		Cmppolicy []Cmppolicy
	}

	if err := c.listResources("cmppolicy", &results); err != nil {
		return nil, err
	}

	return results.Cmppolicy, nil
}

func (c *NitroClient) AddCmppolicy(resource Cmppolicy) error {
	return c.addResource("cmppolicy", resource)
}

func (c *NitroClient) RenameCmppolicy(name string, newName string) error {
	return c.renameResource("cmppolicy", "name", name, newName)
}

func (c *NitroClient) UnsetCmppolicy(name string, fields ...string) error {
	return c.unsetResource("cmppolicy", "name", name, fields)
}

func (c *NitroClient) UpdateCmppolicy(resource Cmppolicy) error {
	update := cmppolicy_update{
		resource.Name,
		resource.Rule,
		resource.Resaction,
	}

	return c.Put("cmppolicy", update)
}
