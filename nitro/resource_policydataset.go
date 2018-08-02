package nitro

type Policydataset struct {
	Name      string `json:"name"`
	Comment   string `json:"comment,omitempty"`
	Indextype string `json:"indextype,omitempty"`
	Type      string `json:"type,omitempty"`
}

type PolicydatasetKey struct {
	Name string
}

type policydataset_update struct {
	Name string `json:"name"`
}

type policydataset_payload struct {
	policydataset interface{}
}

func policydataset_key_to_args(key PolicydatasetKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeletePolicydataset(key PolicydatasetKey) error {
	return c.deleteResourceWithArgs("policydataset", key.Name, policydataset_key_to_args(key))
}

func (c *NitroClient) GetPolicydataset(key PolicydatasetKey) (*Policydataset, error) {
	var results struct {
		Policydataset []Policydataset
	}

	if err := c.getResourceWithArgs("policydataset", key.Name, policydataset_key_to_args(key), &results); err != nil || len(results.Policydataset) != 1 {
		return nil, err
	}

	return &results.Policydataset[0], nil
}

func (c *NitroClient) ListPolicydataset() ([]Policydataset, error) {
	var results struct {
		Policydataset []Policydataset
	}

	if err := c.listResources("policydataset", &results); err != nil {
		return nil, err
	}

	return results.Policydataset, nil
}

func (c *NitroClient) AddPolicydataset(resource Policydataset) error {
	return c.addResource("policydataset", resource)
}

func (c *NitroClient) RenamePolicydataset(name string, newName string) error {
	return c.renameResource("policydataset", "name", name, newName)
}
