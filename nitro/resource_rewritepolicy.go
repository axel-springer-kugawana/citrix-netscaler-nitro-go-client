package nitro

type Rewritepolicy struct {
	Name        string `json:"name"`
	Action      string `json:"action,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
}

type RewritepolicyKey struct {
	Name string
}

type rewritepolicy_update struct {
	Name        string `json:"name"`
	Rule        string `json:"rule,omitempty"`
	Action      string `json:"action,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
}

type rewritepolicy_payload struct {
	rewritepolicy interface{}
}

func rewritepolicy_key_to_args(key RewritepolicyKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteRewritepolicy(key RewritepolicyKey) error {
	return c.deleteResourceWithArgs("rewritepolicy", key.Name, rewritepolicy_key_to_args(key))
}

func (c *NitroClient) GetRewritepolicy(key RewritepolicyKey) (*Rewritepolicy, error) {
	var results struct {
		Rewritepolicy []Rewritepolicy
	}

	if err := c.getResourceWithArgs("rewritepolicy", key.Name, rewritepolicy_key_to_args(key), &results); err != nil || len(results.Rewritepolicy) != 1 {
		return nil, err
	}

	return &results.Rewritepolicy[0], nil
}

func (c *NitroClient) ListRewritepolicy() ([]Rewritepolicy, error) {
	var results struct {
		Rewritepolicy []Rewritepolicy
	}

	if err := c.listResources("rewritepolicy", &results); err != nil {
		return nil, err
	}

	return results.Rewritepolicy, nil
}

func (c *NitroClient) AddRewritepolicy(resource Rewritepolicy) error {
	return c.addResource("rewritepolicy", resource)
}

func (c *NitroClient) RenameRewritepolicy(name string, newName string) error {
	return c.renameResource("rewritepolicy", "name", name, newName)
}

func (c *NitroClient) UnsetRewritepolicy(name string, fields ...string) error {
	return c.unsetResource("rewritepolicy", "name", name, fields)
}

func (c *NitroClient) UpdateRewritepolicy(resource Rewritepolicy) error {
	update := rewritepolicy_update{
		resource.Name,
		resource.Rule,
		resource.Action,
		resource.Undefaction,
		resource.Comment,
		resource.Logaction,
	}

	return c.Put("rewritepolicy", update)
}
