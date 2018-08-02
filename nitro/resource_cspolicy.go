package nitro

type CspolicyResource struct {
	Policyname string `json:"policyname"`
	Action     string `json:"action,omitempty"`
	Domain     string `json:"domain,omitempty"`
	Logaction  string `json:"logaction,omitempty"`
	Rule       string `json:"rule,omitempty"`
	Url        string `json:"url,omitempty"`
}

type CspolicyKey struct {
	Policyname string
}

type cspolicy_update struct {
	Policyname string `json:"policyname"`
	Url        string `json:"url,omitempty"`
	Rule       string `json:"rule,omitempty"`
	Domain     string `json:"domain,omitempty"`
	Action     string `json:"action,omitempty"`
	Logaction  string `json:"logaction,omitempty"`
}

type cspolicy_payload struct {
	cspolicy interface{}
}

func cspolicy_key_to_args(key CspolicyKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteCspolicy(key CspolicyKey) error {
	return c.deleteResourceWithArgs("cspolicy", key.Policyname, cspolicy_key_to_args(key))
}

func (c *NitroClient) GetCspolicy(key CspolicyKey) (*CspolicyResource, error) {
	var results struct {
		Cspolicy []CspolicyResource
	}

	if err := c.getResourceWithArgs("cspolicy", key.Policyname, cspolicy_key_to_args(key), &results); err != nil || len(results.Cspolicy) != 1 {
		return nil, err
	}

	return &results.Cspolicy[0], nil
}

func (c *NitroClient) ListCspolicy() ([]CspolicyResource, error) {
	var results struct {
		Cspolicy []CspolicyResource
	}

	if err := c.listResources("cspolicy", &results); err != nil {
		return nil, err
	}

	return results.Cspolicy, nil
}

func (c *NitroClient) AddCspolicy(resource CspolicyResource) error {
	return c.addResource("cspolicy", resource)
}

func (c *NitroClient) RenameCspolicy(policyname string, newName string) error {
	return c.renameResource("cspolicy", "policyname", policyname, newName)
}

func (c *NitroClient) UnsetCspolicy(policyname string, fields ...string) error {
	return c.unsetResource("cspolicy", "policyname", policyname, fields)
}

func (c *NitroClient) UpdateCspolicy(resource CspolicyResource) error {
	update := cspolicy_update{
		resource.Policyname,
		resource.Url,
		resource.Rule,
		resource.Domain,
		resource.Action,
		resource.Logaction,
	}

	return c.Put("cspolicy", update)
}
