package nitro

type PolicystringmapResource struct {
	Name    string `json:"name"`
	Comment string `json:"comment,omitempty"`
}

type PolicystringmapKey struct {
	Name string
}

type policystringmap_update struct {
	Name    string `json:"name"`
	Comment string `json:"comment,omitempty"`
}

type policystringmap_payload struct {
	policystringmap interface{}
}

func policystringmap_key_to_args(key PolicystringmapKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeletePolicystringmap(key PolicystringmapKey) error {
	return c.deleteResourceWithArgs("policystringmap", key.Name, policystringmap_key_to_args(key))
}

func (c *NitroClient) GetPolicystringmap(key PolicystringmapKey) (*PolicystringmapResource, error) {
	var results struct {
		Policystringmap []PolicystringmapResource
	}

	if err := c.getResourceWithArgs("policystringmap", key.Name, policystringmap_key_to_args(key), &results); err != nil || len(results.Policystringmap) != 1 {
		return nil, err
	}

	return &results.Policystringmap[0], nil
}

func (c *NitroClient) ListPolicystringmap() ([]PolicystringmapResource, error) {
	var results struct {
		Policystringmap []PolicystringmapResource
	}

	if err := c.listResources("policystringmap", &results); err != nil {
		return nil, err
	}

	return results.Policystringmap, nil
}

func (c *NitroClient) AddPolicystringmap(resource PolicystringmapResource) error {
	return c.addResource("policystringmap", resource)
}

func (c *NitroClient) RenamePolicystringmap(name string, newName string) error {
	return c.renameResource("policystringmap", "name", name, newName)
}

func (c *NitroClient) UnsetPolicystringmap(name string, fields ...string) error {
	return c.unsetResource("policystringmap", "name", name, fields)
}

func (c *NitroClient) UpdatePolicystringmap(resource PolicystringmapResource) error {
	update := policystringmap_update{
		resource.Name,
		resource.Comment,
	}

	return c.Put("policystringmap", update)
}
