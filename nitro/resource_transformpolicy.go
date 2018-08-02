package nitro

type TransformpolicyResource struct {
	Name        string `json:"name"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
	Profilename string `json:"profilename,omitempty"`
	Rule        string `json:"rule,omitempty"`
}

type TransformpolicyKey struct {
	Name string
}

type transformpolicy_update struct {
	Name        string `json:"name"`
	Rule        string `json:"rule,omitempty"`
	Profilename string `json:"profilename,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
}

type transformpolicy_payload struct {
	transformpolicy interface{}
}

func transformpolicy_key_to_args(key TransformpolicyKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteTransformpolicy(key TransformpolicyKey) error {
	return c.deleteResourceWithArgs("transformpolicy", key.Name, transformpolicy_key_to_args(key))
}

func (c *NitroClient) GetTransformpolicy(key TransformpolicyKey) (*TransformpolicyResource, error) {
	var results struct {
		Transformpolicy []TransformpolicyResource
	}

	if err := c.getResourceWithArgs("transformpolicy", key.Name, transformpolicy_key_to_args(key), &results); err != nil || len(results.Transformpolicy) != 1 {
		return nil, err
	}

	return &results.Transformpolicy[0], nil
}

func (c *NitroClient) ListTransformpolicy() ([]TransformpolicyResource, error) {
	var results struct {
		Transformpolicy []TransformpolicyResource
	}

	if err := c.listResources("transformpolicy", &results); err != nil {
		return nil, err
	}

	return results.Transformpolicy, nil
}

func (c *NitroClient) AddTransformpolicy(resource TransformpolicyResource) error {
	return c.addResource("transformpolicy", resource)
}

func (c *NitroClient) RenameTransformpolicy(name string, newName string) error {
	return c.renameResource("transformpolicy", "name", name, newName)
}

func (c *NitroClient) UnsetTransformpolicy(name string, fields ...string) error {
	return c.unsetResource("transformpolicy", "name", name, fields)
}

func (c *NitroClient) UpdateTransformpolicy(resource TransformpolicyResource) error {
	update := transformpolicy_update{
		resource.Name,
		resource.Rule,
		resource.Profilename,
		resource.Comment,
		resource.Logaction,
	}

	return c.Put("transformpolicy", update)
}
