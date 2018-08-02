package nitro

type Dnspolicy64Resource struct {
	Name   string `json:"name"`
	Action string `json:"action,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

type Dnspolicy64Key struct {
	Name string
}

type dnspolicy64_update struct {
	Name   string `json:"name"`
	Rule   string `json:"rule,omitempty"`
	Action string `json:"action,omitempty"`
}

type dnspolicy64_payload struct {
	dnspolicy64 interface{}
}

func dnspolicy64_key_to_args(key Dnspolicy64Key) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteDnspolicy64(key Dnspolicy64Key) error {
	return c.deleteResourceWithArgs("dnspolicy64", key.Name, dnspolicy64_key_to_args(key))
}

func (c *NitroClient) GetDnspolicy64(key Dnspolicy64Key) (*Dnspolicy64Resource, error) {
	var results struct {
		Dnspolicy64 []Dnspolicy64Resource
	}

	if err := c.getResourceWithArgs("dnspolicy64", key.Name, dnspolicy64_key_to_args(key), &results); err != nil || len(results.Dnspolicy64) != 1 {
		return nil, err
	}

	return &results.Dnspolicy64[0], nil
}

func (c *NitroClient) ListDnspolicy64() ([]Dnspolicy64Resource, error) {
	var results struct {
		Dnspolicy64 []Dnspolicy64Resource
	}

	if err := c.listResources("dnspolicy64", &results); err != nil {
		return nil, err
	}

	return results.Dnspolicy64, nil
}

func (c *NitroClient) AddDnspolicy64(resource Dnspolicy64Resource) error {
	return c.addResource("dnspolicy64", resource)
}

func (c *NitroClient) RenameDnspolicy64(name string, newName string) error {
	return c.renameResource("dnspolicy64", "name", name, newName)
}

func (c *NitroClient) UnsetDnspolicy64(name string, fields ...string) error {
	return c.unsetResource("dnspolicy64", "name", name, fields)
}

func (c *NitroClient) UpdateDnspolicy64(resource Dnspolicy64Resource) error {
	update := dnspolicy64_update{
		resource.Name,
		resource.Rule,
		resource.Action,
	}

	return c.Put("dnspolicy64", update)
}
