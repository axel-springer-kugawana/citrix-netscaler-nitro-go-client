package nitro

type AppflowpolicyResource struct {
	Name        string `json:"name"`
	Action      string `json:"action,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
}

type AppflowpolicyKey struct {
	Name string
}

type appflowpolicy_update struct {
	Name        string `json:"name"`
	Rule        string `json:"rule,omitempty"`
	Action      string `json:"action,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
	Comment     string `json:"comment,omitempty"`
}

type appflowpolicy_payload struct {
	appflowpolicy interface{}
}

func appflowpolicy_key_to_args(key AppflowpolicyKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteAppflowpolicy(key AppflowpolicyKey) error {
	return c.deleteResourceWithArgs("appflowpolicy", key.Name, appflowpolicy_key_to_args(key))
}

func (c *NitroClient) GetAppflowpolicy(key AppflowpolicyKey) (*AppflowpolicyResource, error) {
	var results struct {
		Appflowpolicy []AppflowpolicyResource
	}

	if err := c.getResourceWithArgs("appflowpolicy", key.Name, appflowpolicy_key_to_args(key), &results); err != nil || len(results.Appflowpolicy) != 1 {
		return nil, err
	}

	return &results.Appflowpolicy[0], nil
}

func (c *NitroClient) ListAppflowpolicy() ([]AppflowpolicyResource, error) {
	var results struct {
		Appflowpolicy []AppflowpolicyResource
	}

	if err := c.listResources("appflowpolicy", &results); err != nil {
		return nil, err
	}

	return results.Appflowpolicy, nil
}

func (c *NitroClient) AddAppflowpolicy(resource AppflowpolicyResource) error {
	return c.addResource("appflowpolicy", resource)
}

func (c *NitroClient) RenameAppflowpolicy(name string, newName string) error {
	return c.renameResource("appflowpolicy", "name", name, newName)
}

func (c *NitroClient) UnsetAppflowpolicy(name string, fields ...string) error {
	return c.unsetResource("appflowpolicy", "name", name, fields)
}

func (c *NitroClient) UpdateAppflowpolicy(resource AppflowpolicyResource) error {
	update := appflowpolicy_update{
		resource.Name,
		resource.Rule,
		resource.Action,
		resource.Undefaction,
		resource.Comment,
	}

	return c.Put("appflowpolicy", update)
}
