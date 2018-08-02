package nitro

type CspolicylabelResource struct {
	Labelname         string `json:"labelname"`
	Cspolicylabeltype string `json:"cspolicylabeltype,omitempty"`
}

type CspolicylabelKey struct {
	Labelname string
}

type cspolicylabel_update struct {
	Labelname string `json:"labelname"`
}

type cspolicylabel_payload struct {
	cspolicylabel interface{}
}

func cspolicylabel_key_to_args(key CspolicylabelKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteCspolicylabel(key CspolicylabelKey) error {
	return c.deleteResourceWithArgs("cspolicylabel", key.Labelname, cspolicylabel_key_to_args(key))
}

func (c *NitroClient) GetCspolicylabel(key CspolicylabelKey) (*CspolicylabelResource, error) {
	var results struct {
		Cspolicylabel []CspolicylabelResource
	}

	if err := c.getResourceWithArgs("cspolicylabel", key.Labelname, cspolicylabel_key_to_args(key), &results); err != nil || len(results.Cspolicylabel) != 1 {
		return nil, err
	}

	return &results.Cspolicylabel[0], nil
}

func (c *NitroClient) ListCspolicylabel() ([]CspolicylabelResource, error) {
	var results struct {
		Cspolicylabel []CspolicylabelResource
	}

	if err := c.listResources("cspolicylabel", &results); err != nil {
		return nil, err
	}

	return results.Cspolicylabel, nil
}

func (c *NitroClient) AddCspolicylabel(resource CspolicylabelResource) error {
	return c.addResource("cspolicylabel", resource)
}

func (c *NitroClient) RenameCspolicylabel(labelname string, newName string) error {
	return c.renameResource("cspolicylabel", "labelname", labelname, newName)
}
