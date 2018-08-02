package nitro

type ScpolicyResource struct {
	Name              string `json:"name"`
	Action            string `json:"action,omitempty"`
	Altcontentpath    string `json:"altcontentpath,omitempty"`
	Altcontentsvcname string `json:"altcontentsvcname,omitempty"`
	Delay             int    `json:"delay,string,omitempty"`
	Maxconn           int    `json:"maxconn,string,omitempty"`
	Rule              string `json:"rule,omitempty"`
	Url               string `json:"url,omitempty"`
}

type ScpolicyKey struct {
	Name string
}

type scpolicy_update struct {
	Name              string `json:"name"`
	Url               string `json:"url,omitempty"`
	Rule              string `json:"rule,omitempty"`
	Delay             int    `json:"delay,string,omitempty"`
	Maxconn           int    `json:"maxconn,string,omitempty"`
	Action            string `json:"action,omitempty"`
	Altcontentsvcname string `json:"altcontentsvcname,omitempty"`
	Altcontentpath    string `json:"altcontentpath,omitempty"`
}

type scpolicy_payload struct {
	scpolicy interface{}
}

func scpolicy_key_to_args(key ScpolicyKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteScpolicy(key ScpolicyKey) error {
	return c.deleteResourceWithArgs("scpolicy", key.Name, scpolicy_key_to_args(key))
}

func (c *NitroClient) GetScpolicy(key ScpolicyKey) (*ScpolicyResource, error) {
	var results struct {
		Scpolicy []ScpolicyResource
	}

	if err := c.getResourceWithArgs("scpolicy", key.Name, scpolicy_key_to_args(key), &results); err != nil || len(results.Scpolicy) != 1 {
		return nil, err
	}

	return &results.Scpolicy[0], nil
}

func (c *NitroClient) ListScpolicy() ([]ScpolicyResource, error) {
	var results struct {
		Scpolicy []ScpolicyResource
	}

	if err := c.listResources("scpolicy", &results); err != nil {
		return nil, err
	}

	return results.Scpolicy, nil
}

func (c *NitroClient) AddScpolicy(resource ScpolicyResource) error {
	return c.addResource("scpolicy", resource)
}

func (c *NitroClient) RenameScpolicy(name string, newName string) error {
	return c.renameResource("scpolicy", "name", name, newName)
}

func (c *NitroClient) UnsetScpolicy(name string, fields ...string) error {
	return c.unsetResource("scpolicy", "name", name, fields)
}

func (c *NitroClient) UpdateScpolicy(resource ScpolicyResource) error {
	update := scpolicy_update{
		resource.Name,
		resource.Url,
		resource.Rule,
		resource.Delay,
		resource.Maxconn,
		resource.Action,
		resource.Altcontentsvcname,
		resource.Altcontentpath,
	}

	return c.Put("scpolicy", update)
}
