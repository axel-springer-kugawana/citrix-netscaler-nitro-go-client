package nitro

type ResponderpolicyResource struct {
	Name          string `json:"name"`
	Action        string `json:"action,omitempty"`
	Appflowaction string `json:"appflowaction,omitempty"`
	Comment       string `json:"comment,omitempty"`
	Logaction     string `json:"logaction,omitempty"`
	Rule          string `json:"rule,omitempty"`
	Undefaction   string `json:"undefaction,omitempty"`
}

type ResponderpolicyKey struct {
	Name string
}

type responderpolicy_update struct {
	Name          string `json:"name"`
	Rule          string `json:"rule,omitempty"`
	Action        string `json:"action,omitempty"`
	Undefaction   string `json:"undefaction,omitempty"`
	Comment       string `json:"comment,omitempty"`
	Logaction     string `json:"logaction,omitempty"`
	Appflowaction string `json:"appflowaction,omitempty"`
}

type responderpolicy_payload struct {
	responderpolicy interface{}
}

func responderpolicy_key_to_args(key ResponderpolicyKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteResponderpolicy(key ResponderpolicyKey) error {
	return c.deleteResourceWithArgs("responderpolicy", key.Name, responderpolicy_key_to_args(key))
}

func (c *NitroClient) GetResponderpolicy(key ResponderpolicyKey) (*ResponderpolicyResource, error) {
	var results struct {
		Responderpolicy []ResponderpolicyResource
	}

	if err := c.getResourceWithArgs("responderpolicy", key.Name, responderpolicy_key_to_args(key), &results); err != nil || len(results.Responderpolicy) != 1 {
		return nil, err
	}

	return &results.Responderpolicy[0], nil
}

func (c *NitroClient) ListResponderpolicy() ([]ResponderpolicyResource, error) {
	var results struct {
		Responderpolicy []ResponderpolicyResource
	}

	if err := c.listResources("responderpolicy", &results); err != nil {
		return nil, err
	}

	return results.Responderpolicy, nil
}

func (c *NitroClient) AddResponderpolicy(resource ResponderpolicyResource) error {
	return c.addResource("responderpolicy", resource)
}

func (c *NitroClient) RenameResponderpolicy(name string, newName string) error {
	return c.renameResource("responderpolicy", "name", name, newName)
}

func (c *NitroClient) UnsetResponderpolicy(name string, fields ...string) error {
	return c.unsetResource("responderpolicy", "name", name, fields)
}

func (c *NitroClient) UpdateResponderpolicy(resource ResponderpolicyResource) error {
	update := responderpolicy_update{
		resource.Name,
		resource.Rule,
		resource.Action,
		resource.Undefaction,
		resource.Comment,
		resource.Logaction,
		resource.Appflowaction,
	}

	return c.Put("responderpolicy", update)
}
