package nitro

type FeopolicyResource struct {
	Name   string `json:"name"`
	Action string `json:"action,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

type FeopolicyKey struct {
	Name string
}

type feopolicy_update struct {
	Name   string `json:"name"`
	Rule   string `json:"rule,omitempty"`
	Action string `json:"action,omitempty"`
}

type feopolicy_payload struct {
	feopolicy interface{}
}

func feopolicy_key_to_args(key FeopolicyKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteFeopolicy(key FeopolicyKey) error {
	return c.deleteResourceWithArgs("feopolicy", key.Name, feopolicy_key_to_args(key))
}

func (c *NitroClient) GetFeopolicy(key FeopolicyKey) (*FeopolicyResource, error) {
	var results struct {
		Feopolicy []FeopolicyResource
	}

	if err := c.getResourceWithArgs("feopolicy", key.Name, feopolicy_key_to_args(key), &results); err != nil || len(results.Feopolicy) != 1 {
		return nil, err
	}

	return &results.Feopolicy[0], nil
}

func (c *NitroClient) ListFeopolicy() ([]FeopolicyResource, error) {
	var results struct {
		Feopolicy []FeopolicyResource
	}

	if err := c.listResources("feopolicy", &results); err != nil {
		return nil, err
	}

	return results.Feopolicy, nil
}

func (c *NitroClient) AddFeopolicy(resource FeopolicyResource) error {
	return c.addResource("feopolicy", resource)
}

func (c *NitroClient) RenameFeopolicy(name string, newName string) error {
	return c.renameResource("feopolicy", "name", name, newName)
}

func (c *NitroClient) UnsetFeopolicy(name string, fields ...string) error {
	return c.unsetResource("feopolicy", "name", name, fields)
}

func (c *NitroClient) UpdateFeopolicy(resource FeopolicyResource) error {
	update := feopolicy_update{
		resource.Name,
		resource.Rule,
		resource.Action,
	}

	return c.Put("feopolicy", update)
}
