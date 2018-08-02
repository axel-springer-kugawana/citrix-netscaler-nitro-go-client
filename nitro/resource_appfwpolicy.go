package nitro

type Appfwpolicy struct {
	Name        string `json:"name"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
	Profilename string `json:"profilename,omitempty"`
	Rule        string `json:"rule,omitempty"`
}

type AppfwpolicyKey struct {
	Name string
}

type appfwpolicy_update struct {
	Name        string `json:"name"`
	Rule        string `json:"rule,omitempty"`
	Profilename string `json:"profilename,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
}

type appfwpolicy_payload struct {
	appfwpolicy interface{}
}

func appfwpolicy_key_to_args(key AppfwpolicyKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteAppfwpolicy(key AppfwpolicyKey) error {
	return c.deleteResourceWithArgs("appfwpolicy", key.Name, appfwpolicy_key_to_args(key))
}

func (c *NitroClient) GetAppfwpolicy(key AppfwpolicyKey) (*Appfwpolicy, error) {
	var results struct {
		Appfwpolicy []Appfwpolicy
	}

	if err := c.getResourceWithArgs("appfwpolicy", key.Name, appfwpolicy_key_to_args(key), &results); err != nil || len(results.Appfwpolicy) != 1 {
		return nil, err
	}

	return &results.Appfwpolicy[0], nil
}

func (c *NitroClient) ListAppfwpolicy() ([]Appfwpolicy, error) {
	var results struct {
		Appfwpolicy []Appfwpolicy
	}

	if err := c.listResources("appfwpolicy", &results); err != nil {
		return nil, err
	}

	return results.Appfwpolicy, nil
}

func (c *NitroClient) AddAppfwpolicy(resource Appfwpolicy) error {
	return c.addResource("appfwpolicy", resource)
}

func (c *NitroClient) RenameAppfwpolicy(name string, newName string) error {
	return c.renameResource("appfwpolicy", "name", name, newName)
}

func (c *NitroClient) UnsetAppfwpolicy(name string, fields ...string) error {
	return c.unsetResource("appfwpolicy", "name", name, fields)
}

func (c *NitroClient) UpdateAppfwpolicy(resource Appfwpolicy) error {
	update := appfwpolicy_update{
		resource.Name,
		resource.Rule,
		resource.Profilename,
		resource.Comment,
		resource.Logaction,
	}

	return c.Put("appfwpolicy", update)
}
