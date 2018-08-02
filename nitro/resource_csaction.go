package nitro

type CsactionResource struct {
	Name              string `json:"name"`
	Comment           string `json:"comment,omitempty"`
	Targetlbvserver   string `json:"targetlbvserver,omitempty"`
	Targetvserver     string `json:"targetvserver,omitempty"`
	Targetvserverexpr string `json:"targetvserverexpr,omitempty"`
}

type CsactionKey struct {
	Name string
}

type csaction_update struct {
	Name              string `json:"name"`
	Targetlbvserver   string `json:"targetlbvserver,omitempty"`
	Targetvserver     string `json:"targetvserver,omitempty"`
	Targetvserverexpr string `json:"targetvserverexpr,omitempty"`
	Comment           string `json:"comment,omitempty"`
}

type csaction_payload struct {
	csaction interface{}
}

func csaction_key_to_args(key CsactionKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteCsaction(key CsactionKey) error {
	return c.deleteResourceWithArgs("csaction", key.Name, csaction_key_to_args(key))
}

func (c *NitroClient) GetCsaction(key CsactionKey) (*CsactionResource, error) {
	var results struct {
		Csaction []CsactionResource
	}

	if err := c.getResourceWithArgs("csaction", key.Name, csaction_key_to_args(key), &results); err != nil || len(results.Csaction) != 1 {
		return nil, err
	}

	return &results.Csaction[0], nil
}

func (c *NitroClient) ListCsaction() ([]CsactionResource, error) {
	var results struct {
		Csaction []CsactionResource
	}

	if err := c.listResources("csaction", &results); err != nil {
		return nil, err
	}

	return results.Csaction, nil
}

func (c *NitroClient) AddCsaction(resource CsactionResource) error {
	return c.addResource("csaction", resource)
}

func (c *NitroClient) RenameCsaction(name string, newName string) error {
	return c.renameResource("csaction", "name", name, newName)
}

func (c *NitroClient) UnsetCsaction(name string, fields ...string) error {
	return c.unsetResource("csaction", "name", name, fields)
}

func (c *NitroClient) UpdateCsaction(resource CsactionResource) error {
	update := csaction_update{
		resource.Name,
		resource.Targetlbvserver,
		resource.Targetvserver,
		resource.Targetvserverexpr,
		resource.Comment,
	}

	return c.Put("csaction", update)
}
