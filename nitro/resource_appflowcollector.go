package nitro

type Appflowcollector struct {
	Name       string `json:"name"`
	Ipaddress  string `json:"ipaddress,omitempty"`
	Netprofile string `json:"netprofile,omitempty"`
	Port       int    `json:"port,omitempty"`
	Transport  string `json:"transport,omitempty"`
}

type AppflowcollectorKey struct {
	Name string
}

type appflowcollector_update struct {
	Name       string `json:"name"`
	Ipaddress  string `json:"ipaddress,omitempty"`
	Port       int    `json:"port,omitempty"`
	Netprofile string `json:"netprofile,omitempty"`
}

type appflowcollector_payload struct {
	appflowcollector interface{}
}

func appflowcollector_key_to_args(key AppflowcollectorKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteAppflowcollector(key AppflowcollectorKey) error {
	return c.deleteResourceWithArgs("appflowcollector", key.Name, appflowcollector_key_to_args(key))
}

func (c *NitroClient) GetAppflowcollector(key AppflowcollectorKey) (*Appflowcollector, error) {
	var results struct {
		Appflowcollector []Appflowcollector
	}

	if err := c.getResourceWithArgs("appflowcollector", key.Name, appflowcollector_key_to_args(key), &results); err != nil || len(results.Appflowcollector) != 1 {
		return nil, err
	}

	return &results.Appflowcollector[0], nil
}

func (c *NitroClient) ListAppflowcollector() ([]Appflowcollector, error) {
	var results struct {
		Appflowcollector []Appflowcollector
	}

	if err := c.listResources("appflowcollector", &results); err != nil {
		return nil, err
	}

	return results.Appflowcollector, nil
}

func (c *NitroClient) AddAppflowcollector(resource Appflowcollector) error {
	return c.addResource("appflowcollector", resource)
}

func (c *NitroClient) RenameAppflowcollector(name string, newName string) error {
	return c.renameResource("appflowcollector", "name", name, newName)
}

func (c *NitroClient) UnsetAppflowcollector(name string, fields ...string) error {
	return c.unsetResource("appflowcollector", "name", name, fields)
}

func (c *NitroClient) UpdateAppflowcollector(resource Appflowcollector) error {
	update := appflowcollector_update{
		resource.Name,
		resource.Ipaddress,
		resource.Port,
		resource.Netprofile,
	}

	return c.Put("appflowcollector", update)
}
