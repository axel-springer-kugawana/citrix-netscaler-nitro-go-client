package nitro

type Dnsaction64 struct {
	Actionname  string `json:"actionname"`
	Excluderule string `json:"excluderule,omitempty"`
	Mappedrule  string `json:"mappedrule,omitempty"`
	Prefix      string `json:"prefix,omitempty"`
}

type Dnsaction64Key struct {
	Actionname string
}

type dnsaction64_update struct {
	Actionname  string `json:"actionname"`
	Prefix      string `json:"prefix,omitempty"`
	Mappedrule  string `json:"mappedrule,omitempty"`
	Excluderule string `json:"excluderule,omitempty"`
}

type dnsaction64_payload struct {
	dnsaction64 interface{}
}

func dnsaction64_key_to_args(key Dnsaction64Key) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteDnsaction64(key Dnsaction64Key) error {
	return c.deleteResourceWithArgs("dnsaction64", key.Actionname, dnsaction64_key_to_args(key))
}

func (c *NitroClient) GetDnsaction64(key Dnsaction64Key) (*Dnsaction64, error) {
	var results struct {
		Dnsaction64 []Dnsaction64
	}

	if err := c.getResourceWithArgs("dnsaction64", key.Actionname, dnsaction64_key_to_args(key), &results); err != nil || len(results.Dnsaction64) != 1 {
		return nil, err
	}

	return &results.Dnsaction64[0], nil
}

func (c *NitroClient) ListDnsaction64() ([]Dnsaction64, error) {
	var results struct {
		Dnsaction64 []Dnsaction64
	}

	if err := c.listResources("dnsaction64", &results); err != nil {
		return nil, err
	}

	return results.Dnsaction64, nil
}

func (c *NitroClient) AddDnsaction64(resource Dnsaction64) error {
	return c.addResource("dnsaction64", resource)
}

func (c *NitroClient) RenameDnsaction64(actionname string, newName string) error {
	return c.renameResource("dnsaction64", "actionname", actionname, newName)
}

func (c *NitroClient) UnsetDnsaction64(actionname string, fields ...string) error {
	return c.unsetResource("dnsaction64", "actionname", actionname, fields)
}

func (c *NitroClient) UpdateDnsaction64(resource Dnsaction64) error {
	update := dnsaction64_update{
		resource.Actionname,
		resource.Prefix,
		resource.Mappedrule,
		resource.Excluderule,
	}

	return c.Put("dnsaction64", update)
}
