package nitro

type Cmppolicylabel struct {
	Labelname string `json:"labelname"`
	Type      string `json:"type,omitempty"`
}

type CmppolicylabelKey struct {
	Labelname string
}

type cmppolicylabel_update struct {
	Labelname string `json:"labelname"`
}

type cmppolicylabel_payload struct {
	cmppolicylabel interface{}
}

func cmppolicylabel_key_to_args(key CmppolicylabelKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteCmppolicylabel(key CmppolicylabelKey) error {
	return c.deleteResourceWithArgs("cmppolicylabel", key.Labelname, cmppolicylabel_key_to_args(key))
}

func (c *NitroClient) GetCmppolicylabel(key CmppolicylabelKey) (*Cmppolicylabel, error) {
	var results struct {
		Cmppolicylabel []Cmppolicylabel
	}

	if err := c.getResourceWithArgs("cmppolicylabel", key.Labelname, cmppolicylabel_key_to_args(key), &results); err != nil || len(results.Cmppolicylabel) != 1 {
		return nil, err
	}

	return &results.Cmppolicylabel[0], nil
}

func (c *NitroClient) ListCmppolicylabel() ([]Cmppolicylabel, error) {
	var results struct {
		Cmppolicylabel []Cmppolicylabel
	}

	if err := c.listResources("cmppolicylabel", &results); err != nil {
		return nil, err
	}

	return results.Cmppolicylabel, nil
}

func (c *NitroClient) AddCmppolicylabel(resource Cmppolicylabel) error {
	return c.addResource("cmppolicylabel", resource)
}

func (c *NitroClient) RenameCmppolicylabel(labelname string, newName string) error {
	return c.renameResource("cmppolicylabel", "labelname", labelname, newName)
}
