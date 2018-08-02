package nitro

type Caaction struct {
	Name         string `json:"name"`
	Accumressize int    `json:"accumressize,string,omitempty"`
	Comment      string `json:"comment,omitempty"`
	Lbvserver    string `json:"lbvserver,omitempty"`
	Type         string `json:"type,omitempty"`
}

type CaactionKey struct {
	Name string
}

type caaction_update struct {
	Name         string `json:"name"`
	Accumressize int    `json:"accumressize,string,omitempty"`
	Lbvserver    string `json:"lbvserver,omitempty"`
	Comment      string `json:"comment,omitempty"`
	Type         string `json:"type,omitempty"`
}

type caaction_payload struct {
	caaction interface{}
}

func caaction_key_to_args(key CaactionKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteCaaction(key CaactionKey) error {
	return c.deleteResourceWithArgs("caaction", key.Name, caaction_key_to_args(key))
}

func (c *NitroClient) GetCaaction(key CaactionKey) (*Caaction, error) {
	var results struct {
		Caaction []Caaction
	}

	if err := c.getResourceWithArgs("caaction", key.Name, caaction_key_to_args(key), &results); err != nil || len(results.Caaction) != 1 {
		return nil, err
	}

	return &results.Caaction[0], nil
}

func (c *NitroClient) ListCaaction() ([]Caaction, error) {
	var results struct {
		Caaction []Caaction
	}

	if err := c.listResources("caaction", &results); err != nil {
		return nil, err
	}

	return results.Caaction, nil
}

func (c *NitroClient) AddCaaction(resource Caaction) error {
	return c.addResource("caaction", resource)
}

func (c *NitroClient) RenameCaaction(name string, newName string) error {
	return c.renameResource("caaction", "name", name, newName)
}

func (c *NitroClient) UnsetCaaction(name string, fields ...string) error {
	return c.unsetResource("caaction", "name", name, fields)
}

func (c *NitroClient) UpdateCaaction(resource Caaction) error {
	update := caaction_update{
		resource.Name,
		resource.Accumressize,
		resource.Lbvserver,
		resource.Comment,
		resource.Type,
	}

	return c.Put("caaction", update)
}
