package nitro

type Tmsessionpolicy struct {
	Name   string `json:"name"`
	Action string `json:"action,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

type TmsessionpolicyKey struct {
	Name string
}

type tmsessionpolicy_update struct {
	Name   string `json:"name"`
	Rule   string `json:"rule,omitempty"`
	Action string `json:"action,omitempty"`
}

type tmsessionpolicy_payload struct {
	tmsessionpolicy interface{}
}

func tmsessionpolicy_key_to_args(key TmsessionpolicyKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteTmsessionpolicy(key TmsessionpolicyKey) error {
	return c.deleteResourceWithArgs("tmsessionpolicy", key.Name, tmsessionpolicy_key_to_args(key))
}

func (c *NitroClient) GetTmsessionpolicy(key TmsessionpolicyKey) (*Tmsessionpolicy, error) {
	var results struct {
		Tmsessionpolicy []Tmsessionpolicy
	}

	if err := c.getResourceWithArgs("tmsessionpolicy", key.Name, tmsessionpolicy_key_to_args(key), &results); err != nil || len(results.Tmsessionpolicy) != 1 {
		return nil, err
	}

	return &results.Tmsessionpolicy[0], nil
}

func (c *NitroClient) ListTmsessionpolicy() ([]Tmsessionpolicy, error) {
	var results struct {
		Tmsessionpolicy []Tmsessionpolicy
	}

	if err := c.listResources("tmsessionpolicy", &results); err != nil {
		return nil, err
	}

	return results.Tmsessionpolicy, nil
}

func (c *NitroClient) AddTmsessionpolicy(resource Tmsessionpolicy) error {
	return c.addResource("tmsessionpolicy", resource)
}

func (c *NitroClient) RenameTmsessionpolicy(name string, newName string) error {
	return c.renameResource("tmsessionpolicy", "name", name, newName)
}

func (c *NitroClient) UnsetTmsessionpolicy(name string, fields ...string) error {
	return c.unsetResource("tmsessionpolicy", "name", name, fields)
}

func (c *NitroClient) UpdateTmsessionpolicy(resource Tmsessionpolicy) error {
	update := tmsessionpolicy_update{
		resource.Name,
		resource.Rule,
		resource.Action,
	}

	return c.Put("tmsessionpolicy", update)
}
