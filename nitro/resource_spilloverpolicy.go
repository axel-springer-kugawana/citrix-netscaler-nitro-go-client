package nitro

type Spilloverpolicy struct {
	Name    string `json:"name"`
	Action  string `json:"action,omitempty"`
	Comment string `json:"comment,omitempty"`
	Rule    string `json:"rule,omitempty"`
}

type SpilloverpolicyKey struct {
	Name string
}

type spilloverpolicy_update struct {
	Name    string `json:"name"`
	Rule    string `json:"rule,omitempty"`
	Action  string `json:"action,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type spilloverpolicy_payload struct {
	spilloverpolicy interface{}
}

func spilloverpolicy_key_to_args(key SpilloverpolicyKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteSpilloverpolicy(key SpilloverpolicyKey) error {
	return c.deleteResourceWithArgs("spilloverpolicy", key.Name, spilloverpolicy_key_to_args(key))
}

func (c *NitroClient) GetSpilloverpolicy(key SpilloverpolicyKey) (*Spilloverpolicy, error) {
	var results struct {
		Spilloverpolicy []Spilloverpolicy
	}

	if err := c.getResourceWithArgs("spilloverpolicy", key.Name, spilloverpolicy_key_to_args(key), &results); err != nil || len(results.Spilloverpolicy) != 1 {
		return nil, err
	}

	return &results.Spilloverpolicy[0], nil
}

func (c *NitroClient) ListSpilloverpolicy() ([]Spilloverpolicy, error) {
	var results struct {
		Spilloverpolicy []Spilloverpolicy
	}

	if err := c.listResources("spilloverpolicy", &results); err != nil {
		return nil, err
	}

	return results.Spilloverpolicy, nil
}

func (c *NitroClient) AddSpilloverpolicy(resource Spilloverpolicy) error {
	return c.addResource("spilloverpolicy", resource)
}

func (c *NitroClient) RenameSpilloverpolicy(name string, newName string) error {
	return c.renameResource("spilloverpolicy", "name", name, newName)
}

func (c *NitroClient) UnsetSpilloverpolicy(name string, fields ...string) error {
	return c.unsetResource("spilloverpolicy", "name", name, fields)
}

func (c *NitroClient) UpdateSpilloverpolicy(resource Spilloverpolicy) error {
	update := spilloverpolicy_update{
		resource.Name,
		resource.Rule,
		resource.Action,
		resource.Comment,
	}

	return c.Put("spilloverpolicy", update)
}
