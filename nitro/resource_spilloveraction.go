package nitro

type Spilloveraction struct {
	Name   string `json:"name"`
	Action string `json:"action,omitempty"`
}

type SpilloveractionKey struct {
	Name string
}

type spilloveraction_update struct {
	Name string `json:"name"`
}

type spilloveraction_payload struct {
	spilloveraction interface{}
}

func spilloveraction_key_to_args(key SpilloveractionKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteSpilloveraction(key SpilloveractionKey) error {
	return c.deleteResourceWithArgs("spilloveraction", key.Name, spilloveraction_key_to_args(key))
}

func (c *NitroClient) GetSpilloveraction(key SpilloveractionKey) (*Spilloveraction, error) {
	var results struct {
		Spilloveraction []Spilloveraction
	}

	if err := c.getResourceWithArgs("spilloveraction", key.Name, spilloveraction_key_to_args(key), &results); err != nil || len(results.Spilloveraction) != 1 {
		return nil, err
	}

	return &results.Spilloveraction[0], nil
}

func (c *NitroClient) ListSpilloveraction() ([]Spilloveraction, error) {
	var results struct {
		Spilloveraction []Spilloveraction
	}

	if err := c.listResources("spilloveraction", &results); err != nil {
		return nil, err
	}

	return results.Spilloveraction, nil
}

func (c *NitroClient) AddSpilloveraction(resource Spilloveraction) error {
	return c.addResource("spilloveraction", resource)
}

func (c *NitroClient) RenameSpilloveraction(name string, newName string) error {
	return c.renameResource("spilloveraction", "name", name, newName)
}
