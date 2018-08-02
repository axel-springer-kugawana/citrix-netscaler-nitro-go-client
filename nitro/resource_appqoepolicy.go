package nitro

type AppqoepolicyResource struct {
	Name   string `json:"name"`
	Action string `json:"action,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

type AppqoepolicyKey struct {
	Name string
}

type appqoepolicy_update struct {
	Name   string `json:"name"`
	Rule   string `json:"rule,omitempty"`
	Action string `json:"action,omitempty"`
}

type appqoepolicy_payload struct {
	appqoepolicy interface{}
}

func appqoepolicy_key_to_args(key AppqoepolicyKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteAppqoepolicy(key AppqoepolicyKey) error {
	return c.deleteResourceWithArgs("appqoepolicy", key.Name, appqoepolicy_key_to_args(key))
}

func (c *NitroClient) GetAppqoepolicy(key AppqoepolicyKey) (*AppqoepolicyResource, error) {
	var results struct {
		Appqoepolicy []AppqoepolicyResource
	}

	if err := c.getResourceWithArgs("appqoepolicy", key.Name, appqoepolicy_key_to_args(key), &results); err != nil || len(results.Appqoepolicy) != 1 {
		return nil, err
	}

	return &results.Appqoepolicy[0], nil
}

func (c *NitroClient) ListAppqoepolicy() ([]AppqoepolicyResource, error) {
	var results struct {
		Appqoepolicy []AppqoepolicyResource
	}

	if err := c.listResources("appqoepolicy", &results); err != nil {
		return nil, err
	}

	return results.Appqoepolicy, nil
}

func (c *NitroClient) AddAppqoepolicy(resource AppqoepolicyResource) error {
	return c.addResource("appqoepolicy", resource)
}

func (c *NitroClient) RenameAppqoepolicy(name string, newName string) error {
	return c.renameResource("appqoepolicy", "name", name, newName)
}

func (c *NitroClient) UnsetAppqoepolicy(name string, fields ...string) error {
	return c.unsetResource("appqoepolicy", "name", name, fields)
}

func (c *NitroClient) UpdateAppqoepolicy(resource AppqoepolicyResource) error {
	update := appqoepolicy_update{
		resource.Name,
		resource.Rule,
		resource.Action,
	}

	return c.Put("appqoepolicy", update)
}
