package nitro

type NetprofileResource struct {
	Name             string `json:"name"`
	Overridelsn      string `json:"overridelsn,omitempty"`
	Srcip            string `json:"srcip,omitempty"`
	Srcippersistency string `json:"srcippersistency,omitempty"`
	Td               int    `json:"td,string,omitempty"`
}

type NetprofileKey struct {
	Name string
}

type netprofile_update struct {
	Name             string `json:"name"`
	Srcip            string `json:"srcip,omitempty"`
	Srcippersistency string `json:"srcippersistency,omitempty"`
	Overridelsn      string `json:"overridelsn,omitempty"`
}

type netprofile_payload struct {
	netprofile interface{}
}

func netprofile_key_to_args(key NetprofileKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteNetprofile(key NetprofileKey) error {
	return c.deleteResourceWithArgs("netprofile", key.Name, netprofile_key_to_args(key))
}

func (c *NitroClient) GetNetprofile(key NetprofileKey) (*NetprofileResource, error) {
	var results struct {
		Netprofile []NetprofileResource
	}

	if err := c.getResourceWithArgs("netprofile", key.Name, netprofile_key_to_args(key), &results); err != nil || len(results.Netprofile) != 1 {
		return nil, err
	}

	return &results.Netprofile[0], nil
}

func (c *NitroClient) ListNetprofile() ([]NetprofileResource, error) {
	var results struct {
		Netprofile []NetprofileResource
	}

	if err := c.listResources("netprofile", &results); err != nil {
		return nil, err
	}

	return results.Netprofile, nil
}

func (c *NitroClient) AddNetprofile(resource NetprofileResource) error {
	return c.addResource("netprofile", resource)
}

func (c *NitroClient) RenameNetprofile(name string, newName string) error {
	return c.renameResource("netprofile", "name", name, newName)
}

func (c *NitroClient) UnsetNetprofile(name string, fields ...string) error {
	return c.unsetResource("netprofile", "name", name, fields)
}

func (c *NitroClient) UpdateNetprofile(resource NetprofileResource) error {
	update := netprofile_update{
		resource.Name,
		resource.Srcip,
		resource.Srcippersistency,
		resource.Overridelsn,
	}

	return c.Put("netprofile", update)
}
