package nitro

type Lbmetrictable struct {
	Metrictable string `json:"metrictable"`
}

type LbmetrictableKey struct {
	Metrictable string
}

type lbmetrictable_update struct {
	Metrictable string `json:"metrictable"`
}

type lbmetrictable_payload struct {
	lbmetrictable interface{}
}

func lbmetrictable_key_to_args(key LbmetrictableKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteLbmetrictable(key LbmetrictableKey) error {
	return c.deleteResourceWithArgs("lbmetrictable", key.Metrictable, lbmetrictable_key_to_args(key))
}

func (c *NitroClient) GetLbmetrictable(key LbmetrictableKey) (*Lbmetrictable, error) {
	var results struct {
		Lbmetrictable []Lbmetrictable
	}

	if err := c.getResourceWithArgs("lbmetrictable", key.Metrictable, lbmetrictable_key_to_args(key), &results); err != nil || len(results.Lbmetrictable) != 1 {
		return nil, err
	}

	return &results.Lbmetrictable[0], nil
}

func (c *NitroClient) ListLbmetrictable() ([]Lbmetrictable, error) {
	var results struct {
		Lbmetrictable []Lbmetrictable
	}

	if err := c.listResources("lbmetrictable", &results); err != nil {
		return nil, err
	}

	return results.Lbmetrictable, nil
}

func (c *NitroClient) AddLbmetrictable(resource Lbmetrictable) error {
	return c.addResource("lbmetrictable", resource)
}

func (c *NitroClient) RenameLbmetrictable(metrictable string, newName string) error {
	return c.renameResource("lbmetrictable", "metrictable", metrictable, newName)
}
