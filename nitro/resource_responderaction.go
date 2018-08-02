package nitro

type ResponderactionResource struct {
	Name               string `json:"name"`
	Bypasssafetycheck  string `json:"bypasssafetycheck,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Htmlpage           string `json:"htmlpage,omitempty"`
	Reasonphrase       string `json:"reasonphrase,omitempty"`
	Responsestatuscode int    `json:"responsestatuscode,string,omitempty"`
	Target             string `json:"target,omitempty"`
	Type               string `json:"type,omitempty"`
}

type ResponderactionKey struct {
	Name string
}

type responderaction_update struct {
	Name               string `json:"name"`
	Target             string `json:"target,omitempty"`
	Bypasssafetycheck  string `json:"bypasssafetycheck,omitempty"`
	Htmlpage           string `json:"htmlpage,omitempty"`
	Responsestatuscode int    `json:"responsestatuscode,string,omitempty"`
	Reasonphrase       string `json:"reasonphrase,omitempty"`
	Comment            string `json:"comment,omitempty"`
}

type responderaction_payload struct {
	responderaction interface{}
}

func responderaction_key_to_args(key ResponderactionKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteResponderaction(key ResponderactionKey) error {
	return c.deleteResourceWithArgs("responderaction", key.Name, responderaction_key_to_args(key))
}

func (c *NitroClient) GetResponderaction(key ResponderactionKey) (*ResponderactionResource, error) {
	var results struct {
		Responderaction []ResponderactionResource
	}

	if err := c.getResourceWithArgs("responderaction", key.Name, responderaction_key_to_args(key), &results); err != nil || len(results.Responderaction) != 1 {
		return nil, err
	}

	return &results.Responderaction[0], nil
}

func (c *NitroClient) ListResponderaction() ([]ResponderactionResource, error) {
	var results struct {
		Responderaction []ResponderactionResource
	}

	if err := c.listResources("responderaction", &results); err != nil {
		return nil, err
	}

	return results.Responderaction, nil
}

func (c *NitroClient) AddResponderaction(resource ResponderactionResource) error {
	return c.addResource("responderaction", resource)
}

func (c *NitroClient) RenameResponderaction(name string, newName string) error {
	return c.renameResource("responderaction", "name", name, newName)
}

func (c *NitroClient) UnsetResponderaction(name string, fields ...string) error {
	return c.unsetResource("responderaction", "name", name, fields)
}

func (c *NitroClient) UpdateResponderaction(resource ResponderactionResource) error {
	update := responderaction_update{
		resource.Name,
		resource.Target,
		resource.Bypasssafetycheck,
		resource.Htmlpage,
		resource.Responsestatuscode,
		resource.Reasonphrase,
		resource.Comment,
	}

	return c.Put("responderaction", update)
}
