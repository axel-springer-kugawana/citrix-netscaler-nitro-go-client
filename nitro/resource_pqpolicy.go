package nitro

type Pqpolicy struct {
	Policyname string `json:"policyname"`
	Polqdepth  int    `json:"polqdepth,string,omitempty"`
	Priority   int    `json:"priority,string,omitempty"`
	Qdepth     int    `json:"qdepth,string,omitempty"`
	Rule       string `json:"rule,omitempty"`
	Weight     int    `json:"weight,string,omitempty"`
}

type PqpolicyKey struct {
	Policyname string
}

type pqpolicy_update struct {
	Policyname string `json:"policyname"`
	Weight     int    `json:"weight,string,omitempty"`
	Qdepth     int    `json:"qdepth,string,omitempty"`
	Polqdepth  int    `json:"polqdepth,string,omitempty"`
}

type pqpolicy_payload struct {
	pqpolicy interface{}
}

func pqpolicy_key_to_args(key PqpolicyKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeletePqpolicy(key PqpolicyKey) error {
	return c.deleteResourceWithArgs("pqpolicy", key.Policyname, pqpolicy_key_to_args(key))
}

func (c *NitroClient) GetPqpolicy(key PqpolicyKey) (*Pqpolicy, error) {
	var results struct {
		Pqpolicy []Pqpolicy
	}

	if err := c.getResourceWithArgs("pqpolicy", key.Policyname, pqpolicy_key_to_args(key), &results); err != nil || len(results.Pqpolicy) != 1 {
		return nil, err
	}

	return &results.Pqpolicy[0], nil
}

func (c *NitroClient) ListPqpolicy() ([]Pqpolicy, error) {
	var results struct {
		Pqpolicy []Pqpolicy
	}

	if err := c.listResources("pqpolicy", &results); err != nil {
		return nil, err
	}

	return results.Pqpolicy, nil
}

func (c *NitroClient) AddPqpolicy(resource Pqpolicy) error {
	return c.addResource("pqpolicy", resource)
}

func (c *NitroClient) RenamePqpolicy(policyname string, newName string) error {
	return c.renameResource("pqpolicy", "policyname", policyname, newName)
}

func (c *NitroClient) UnsetPqpolicy(policyname string, fields ...string) error {
	return c.unsetResource("pqpolicy", "policyname", policyname, fields)
}

func (c *NitroClient) UpdatePqpolicy(resource Pqpolicy) error {
	update := pqpolicy_update{
		resource.Policyname,
		resource.Weight,
		resource.Qdepth,
		resource.Polqdepth,
	}

	return c.Put("pqpolicy", update)
}
