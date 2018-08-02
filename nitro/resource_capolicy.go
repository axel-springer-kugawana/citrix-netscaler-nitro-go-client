package nitro

type Capolicy struct {
	Name        string `json:"name"`
	Action      string `json:"action,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
}

type CapolicyKey struct {
	Name string
}

type capolicy_update struct {
	Name        string `json:"name"`
	Rule        string `json:"rule,omitempty"`
	Action      string `json:"action,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
}

type capolicy_payload struct {
	capolicy interface{}
}

func capolicy_key_to_args(key CapolicyKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteCapolicy(key CapolicyKey) error {
	return c.deleteResourceWithArgs("capolicy", key.Name, capolicy_key_to_args(key))
}

func (c *NitroClient) GetCapolicy(key CapolicyKey) (*Capolicy, error) {
	var results struct {
		Capolicy []Capolicy
	}

	if err := c.getResourceWithArgs("capolicy", key.Name, capolicy_key_to_args(key), &results); err != nil || len(results.Capolicy) != 1 {
		return nil, err
	}

	return &results.Capolicy[0], nil
}

func (c *NitroClient) ListCapolicy() ([]Capolicy, error) {
	var results struct {
		Capolicy []Capolicy
	}

	if err := c.listResources("capolicy", &results); err != nil {
		return nil, err
	}

	return results.Capolicy, nil
}

func (c *NitroClient) AddCapolicy(resource Capolicy) error {
	return c.addResource("capolicy", resource)
}

func (c *NitroClient) RenameCapolicy(name string, newName string) error {
	return c.renameResource("capolicy", "name", name, newName)
}

func (c *NitroClient) UnsetCapolicy(name string, fields ...string) error {
	return c.unsetResource("capolicy", "name", name, fields)
}

func (c *NitroClient) UpdateCapolicy(resource Capolicy) error {
	update := capolicy_update{
		resource.Name,
		resource.Rule,
		resource.Action,
		resource.Comment,
		resource.Logaction,
		resource.Undefaction,
	}

	return c.Put("capolicy", update)
}
