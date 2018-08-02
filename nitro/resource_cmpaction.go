package nitro

type CmpactionResource struct {
	Name            string `json:"name"`
	Addvaryheader   string `json:"addvaryheader,omitempty"`
	Cmptype         string `json:"cmptype,omitempty"`
	Deltatype       string `json:"deltatype,omitempty"`
	Varyheadervalue string `json:"varyheadervalue,omitempty"`
}

type CmpactionKey struct {
	Name string
}

type cmpaction_update struct {
	Name            string `json:"name"`
	Cmptype         string `json:"cmptype,omitempty"`
	Addvaryheader   string `json:"addvaryheader,omitempty"`
	Varyheadervalue string `json:"varyheadervalue,omitempty"`
}

type cmpaction_payload struct {
	cmpaction interface{}
}

func cmpaction_key_to_args(key CmpactionKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteCmpaction(key CmpactionKey) error {
	return c.deleteResourceWithArgs("cmpaction", key.Name, cmpaction_key_to_args(key))
}

func (c *NitroClient) GetCmpaction(key CmpactionKey) (*CmpactionResource, error) {
	var results struct {
		Cmpaction []CmpactionResource
	}

	if err := c.getResourceWithArgs("cmpaction", key.Name, cmpaction_key_to_args(key), &results); err != nil || len(results.Cmpaction) != 1 {
		return nil, err
	}

	return &results.Cmpaction[0], nil
}

func (c *NitroClient) ListCmpaction() ([]CmpactionResource, error) {
	var results struct {
		Cmpaction []CmpactionResource
	}

	if err := c.listResources("cmpaction", &results); err != nil {
		return nil, err
	}

	return results.Cmpaction, nil
}

func (c *NitroClient) AddCmpaction(resource CmpactionResource) error {
	return c.addResource("cmpaction", resource)
}

func (c *NitroClient) RenameCmpaction(name string, newName string) error {
	return c.renameResource("cmpaction", "name", name, newName)
}

func (c *NitroClient) UnsetCmpaction(name string, fields ...string) error {
	return c.unsetResource("cmpaction", "name", name, fields)
}

func (c *NitroClient) UpdateCmpaction(resource CmpactionResource) error {
	update := cmpaction_update{
		resource.Name,
		resource.Cmptype,
		resource.Addvaryheader,
		resource.Varyheadervalue,
	}

	return c.Put("cmpaction", update)
}
