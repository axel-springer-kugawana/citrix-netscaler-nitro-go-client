package nitro

type Policyexpression struct {
	Name                  string `json:"name"`
	Clientsecuritymessage string `json:"clientsecuritymessage,omitempty"`
	Comment               string `json:"comment,omitempty"`
	Description           string `json:"description,omitempty"`
	Value                 string `json:"value,omitempty"`
}

type PolicyexpressionKey struct {
	Name string
}

type policyexpression_update struct {
	Name                  string `json:"name"`
	Value                 string `json:"value,omitempty"`
	Description           string `json:"description,omitempty"`
	Comment               string `json:"comment,omitempty"`
	Clientsecuritymessage string `json:"clientsecuritymessage,omitempty"`
}

type policyexpression_payload struct {
	policyexpression interface{}
}

func policyexpression_key_to_args(key PolicyexpressionKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeletePolicyexpression(key PolicyexpressionKey) error {
	return c.deleteResourceWithArgs("policyexpression", key.Name, policyexpression_key_to_args(key))
}

func (c *NitroClient) GetPolicyexpression(key PolicyexpressionKey) (*Policyexpression, error) {
	var results struct {
		Policyexpression []Policyexpression
	}

	if err := c.getResourceWithArgs("policyexpression", key.Name, policyexpression_key_to_args(key), &results); err != nil || len(results.Policyexpression) != 1 {
		return nil, err
	}

	return &results.Policyexpression[0], nil
}

func (c *NitroClient) ListPolicyexpression() ([]Policyexpression, error) {
	var results struct {
		Policyexpression []Policyexpression
	}

	if err := c.listResources("policyexpression", &results); err != nil {
		return nil, err
	}

	return results.Policyexpression, nil
}

func (c *NitroClient) AddPolicyexpression(resource Policyexpression) error {
	return c.addResource("policyexpression", resource)
}

func (c *NitroClient) RenamePolicyexpression(name string, newName string) error {
	return c.renameResource("policyexpression", "name", name, newName)
}

func (c *NitroClient) UnsetPolicyexpression(name string, fields ...string) error {
	return c.unsetResource("policyexpression", "name", name, fields)
}

func (c *NitroClient) UpdatePolicyexpression(resource Policyexpression) error {
	update := policyexpression_update{
		resource.Name,
		resource.Value,
		resource.Description,
		resource.Comment,
		resource.Clientsecuritymessage,
	}

	return c.Put("policyexpression", update)
}
