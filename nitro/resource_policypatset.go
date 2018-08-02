package nitro

type Policypatset struct {
	Name      string `json:"name"`
	Comment   string `json:"comment,omitempty"`
	Indextype string `json:"indextype,omitempty"`
}

type PolicypatsetKey struct {
	Name string
}

type policypatset_update struct {
	Name string `json:"name"`
}

type policypatset_payload struct {
	policypatset interface{}
}

func policypatset_key_to_args(key PolicypatsetKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeletePolicypatset(key PolicypatsetKey) error {
	return c.deleteResourceWithArgs("policypatset", key.Name, policypatset_key_to_args(key))
}

func (c *NitroClient) GetPolicypatset(key PolicypatsetKey) (*Policypatset, error) {
	var results struct {
		Policypatset []Policypatset
	}

	if err := c.getResourceWithArgs("policypatset", key.Name, policypatset_key_to_args(key), &results); err != nil || len(results.Policypatset) != 1 {
		return nil, err
	}

	return &results.Policypatset[0], nil
}

func (c *NitroClient) ListPolicypatset() ([]Policypatset, error) {
	var results struct {
		Policypatset []Policypatset
	}

	if err := c.listResources("policypatset", &results); err != nil {
		return nil, err
	}

	return results.Policypatset, nil
}

func (c *NitroClient) AddPolicypatset(resource Policypatset) error {
	return c.addResource("policypatset", resource)
}

func (c *NitroClient) RenamePolicypatset(name string, newName string) error {
	return c.renameResource("policypatset", "name", name, newName)
}
