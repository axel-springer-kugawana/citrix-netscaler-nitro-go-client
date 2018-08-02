package nitro

type LbwlmResource struct {
	Wlmname   string `json:"wlmname"`
	Ipaddress string `json:"ipaddress,omitempty"`
	Katimeout int    `json:"katimeout,string,omitempty"`
	Lbuid     string `json:"lbuid,omitempty"`
	Port      int    `json:"port,omitempty"`
}

type LbwlmKey struct {
	Wlmname string
}

type lbwlm_update struct {
	Wlmname   string `json:"wlmname"`
	Katimeout int    `json:"katimeout,string,omitempty"`
}

type lbwlm_payload struct {
	lbwlm interface{}
}

func lbwlm_key_to_args(key LbwlmKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteLbwlm(key LbwlmKey) error {
	return c.deleteResourceWithArgs("lbwlm", key.Wlmname, lbwlm_key_to_args(key))
}

func (c *NitroClient) GetLbwlm(key LbwlmKey) (*LbwlmResource, error) {
	var results struct {
		Lbwlm []LbwlmResource
	}

	if err := c.getResourceWithArgs("lbwlm", key.Wlmname, lbwlm_key_to_args(key), &results); err != nil || len(results.Lbwlm) != 1 {
		return nil, err
	}

	return &results.Lbwlm[0], nil
}

func (c *NitroClient) ListLbwlm() ([]LbwlmResource, error) {
	var results struct {
		Lbwlm []LbwlmResource
	}

	if err := c.listResources("lbwlm", &results); err != nil {
		return nil, err
	}

	return results.Lbwlm, nil
}

func (c *NitroClient) AddLbwlm(resource LbwlmResource) error {
	return c.addResource("lbwlm", resource)
}

func (c *NitroClient) RenameLbwlm(wlmname string, newName string) error {
	return c.renameResource("lbwlm", "wlmname", wlmname, newName)
}

func (c *NitroClient) UnsetLbwlm(wlmname string, fields ...string) error {
	return c.unsetResource("lbwlm", "wlmname", wlmname, fields)
}

func (c *NitroClient) UpdateLbwlm(resource LbwlmResource) error {
	update := lbwlm_update{
		resource.Wlmname,
		resource.Katimeout,
	}

	return c.Put("lbwlm", update)
}
