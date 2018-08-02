package nitro

type FilterpolicyResource struct {
	Name      string `json:"name"`
	Reqaction string `json:"reqaction,omitempty"`
	Resaction string `json:"resaction,omitempty"`
	Rule      string `json:"rule,omitempty"`
}

type FilterpolicyKey struct {
	Name string
}

type filterpolicy_update struct {
	Name      string `json:"name"`
	Rule      string `json:"rule,omitempty"`
	Reqaction string `json:"reqaction,omitempty"`
	Resaction string `json:"resaction,omitempty"`
}

type filterpolicy_payload struct {
	filterpolicy interface{}
}

func filterpolicy_key_to_args(key FilterpolicyKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteFilterpolicy(key FilterpolicyKey) error {
	return c.deleteResourceWithArgs("filterpolicy", key.Name, filterpolicy_key_to_args(key))
}

func (c *NitroClient) GetFilterpolicy(key FilterpolicyKey) (*FilterpolicyResource, error) {
	var results struct {
		Filterpolicy []FilterpolicyResource
	}

	if err := c.getResourceWithArgs("filterpolicy", key.Name, filterpolicy_key_to_args(key), &results); err != nil || len(results.Filterpolicy) != 1 {
		return nil, err
	}

	return &results.Filterpolicy[0], nil
}

func (c *NitroClient) ListFilterpolicy() ([]FilterpolicyResource, error) {
	var results struct {
		Filterpolicy []FilterpolicyResource
	}

	if err := c.listResources("filterpolicy", &results); err != nil {
		return nil, err
	}

	return results.Filterpolicy, nil
}

func (c *NitroClient) AddFilterpolicy(resource FilterpolicyResource) error {
	return c.addResource("filterpolicy", resource)
}

func (c *NitroClient) RenameFilterpolicy(name string, newName string) error {
	return c.renameResource("filterpolicy", "name", name, newName)
}

func (c *NitroClient) UnsetFilterpolicy(name string, fields ...string) error {
	return c.unsetResource("filterpolicy", "name", name, fields)
}

func (c *NitroClient) UpdateFilterpolicy(resource FilterpolicyResource) error {
	update := filterpolicy_update{
		resource.Name,
		resource.Rule,
		resource.Reqaction,
		resource.Resaction,
	}

	return c.Put("filterpolicy", update)
}
