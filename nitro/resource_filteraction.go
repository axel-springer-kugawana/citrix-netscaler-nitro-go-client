package nitro

type Filteraction struct {
	Name        string `json:"name"`
	Page        string `json:"page,omitempty"`
	Qual        string `json:"qual,omitempty"`
	Respcode    int    `json:"respcode,string,omitempty"`
	Servicename string `json:"servicename,omitempty"`
	Value       string `json:"value,omitempty"`
}

type FilteractionKey struct {
	Name string
}

type filteraction_update struct {
	Name        string `json:"name"`
	Servicename string `json:"servicename,omitempty"`
	Value       string `json:"value,omitempty"`
	Respcode    int    `json:"respcode,string,omitempty"`
	Page        string `json:"page,omitempty"`
}

type filteraction_payload struct {
	filteraction interface{}
}

func filteraction_key_to_args(key FilteractionKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteFilteraction(key FilteractionKey) error {
	return c.deleteResourceWithArgs("filteraction", key.Name, filteraction_key_to_args(key))
}

func (c *NitroClient) GetFilteraction(key FilteractionKey) (*Filteraction, error) {
	var results struct {
		Filteraction []Filteraction
	}

	if err := c.getResourceWithArgs("filteraction", key.Name, filteraction_key_to_args(key), &results); err != nil || len(results.Filteraction) != 1 {
		return nil, err
	}

	return &results.Filteraction[0], nil
}

func (c *NitroClient) ListFilteraction() ([]Filteraction, error) {
	var results struct {
		Filteraction []Filteraction
	}

	if err := c.listResources("filteraction", &results); err != nil {
		return nil, err
	}

	return results.Filteraction, nil
}

func (c *NitroClient) AddFilteraction(resource Filteraction) error {
	return c.addResource("filteraction", resource)
}

func (c *NitroClient) RenameFilteraction(name string, newName string) error {
	return c.renameResource("filteraction", "name", name, newName)
}

func (c *NitroClient) UnsetFilteraction(name string, fields ...string) error {
	return c.unsetResource("filteraction", "name", name, fields)
}

func (c *NitroClient) UpdateFilteraction(resource Filteraction) error {
	update := filteraction_update{
		resource.Name,
		resource.Servicename,
		resource.Value,
		resource.Respcode,
		resource.Page,
	}

	return c.Put("filteraction", update)
}
