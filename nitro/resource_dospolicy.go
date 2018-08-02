package nitro

type Dospolicy struct {
	Name          string `json:"name"`
	Cltdetectrate int    `json:"cltdetectrate,string,omitempty"`
	Qdepth        int    `json:"qdepth,string,omitempty"`
}

type DospolicyKey struct {
	Name string
}

type dospolicy_update struct {
	Name          string `json:"name"`
	Qdepth        int    `json:"qdepth,string,omitempty"`
	Cltdetectrate int    `json:"cltdetectrate,string,omitempty"`
}

type dospolicy_payload struct {
	dospolicy interface{}
}

func dospolicy_key_to_args(key DospolicyKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteDospolicy(key DospolicyKey) error {
	return c.deleteResourceWithArgs("dospolicy", key.Name, dospolicy_key_to_args(key))
}

func (c *NitroClient) GetDospolicy(key DospolicyKey) (*Dospolicy, error) {
	var results struct {
		Dospolicy []Dospolicy
	}

	if err := c.getResourceWithArgs("dospolicy", key.Name, dospolicy_key_to_args(key), &results); err != nil || len(results.Dospolicy) != 1 {
		return nil, err
	}

	return &results.Dospolicy[0], nil
}

func (c *NitroClient) ListDospolicy() ([]Dospolicy, error) {
	var results struct {
		Dospolicy []Dospolicy
	}

	if err := c.listResources("dospolicy", &results); err != nil {
		return nil, err
	}

	return results.Dospolicy, nil
}

func (c *NitroClient) AddDospolicy(resource Dospolicy) error {
	return c.addResource("dospolicy", resource)
}

func (c *NitroClient) RenameDospolicy(name string, newName string) error {
	return c.renameResource("dospolicy", "name", name, newName)
}

func (c *NitroClient) UnsetDospolicy(name string, fields ...string) error {
	return c.unsetResource("dospolicy", "name", name, fields)
}

func (c *NitroClient) UpdateDospolicy(resource Dospolicy) error {
	update := dospolicy_update{
		resource.Name,
		resource.Qdepth,
		resource.Cltdetectrate,
	}

	return c.Put("dospolicy", update)
}
