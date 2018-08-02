package nitro

type AppqoeactionResource struct {
	Name              string `json:"name"`
	Altcontentpath    string `json:"altcontentpath,omitempty"`
	Altcontentsvcname string `json:"altcontentsvcname,omitempty"`
	Customfile        string `json:"customfile,omitempty"`
	Delay             int    `json:"delay,string,omitempty"`
	Dosaction         string `json:"dosaction,omitempty"`
	Dostrigexpression string `json:"dostrigexpression,omitempty"`
	Maxconn           int    `json:"maxconn,string,omitempty"`
	Polqdepth         int    `json:"polqdepth,string,omitempty"`
	Priority          string `json:"priority,omitempty"`
	Priqdepth         int    `json:"priqdepth,string,omitempty"`
	Respondwith       string `json:"respondwith,omitempty"`
	Tcpprofile        string `json:"tcpprofile,omitempty"`
}

type AppqoeactionKey struct {
	Name string
}

type appqoeaction_update struct {
	Name              string `json:"name"`
	Priority          string `json:"priority,omitempty"`
	Altcontentsvcname string `json:"altcontentsvcname,omitempty"`
	Altcontentpath    string `json:"altcontentpath,omitempty"`
	Polqdepth         int    `json:"polqdepth,string,omitempty"`
	Priqdepth         int    `json:"priqdepth,string,omitempty"`
	Maxconn           int    `json:"maxconn,string,omitempty"`
	Delay             int    `json:"delay,string,omitempty"`
	Dostrigexpression string `json:"dostrigexpression,omitempty"`
	Dosaction         string `json:"dosaction,omitempty"`
	Tcpprofile        string `json:"tcpprofile,omitempty"`
}

type appqoeaction_payload struct {
	appqoeaction interface{}
}

func appqoeaction_key_to_args(key AppqoeactionKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteAppqoeaction(key AppqoeactionKey) error {
	return c.deleteResourceWithArgs("appqoeaction", key.Name, appqoeaction_key_to_args(key))
}

func (c *NitroClient) GetAppqoeaction(key AppqoeactionKey) (*AppqoeactionResource, error) {
	var results struct {
		Appqoeaction []AppqoeactionResource
	}

	if err := c.getResourceWithArgs("appqoeaction", key.Name, appqoeaction_key_to_args(key), &results); err != nil || len(results.Appqoeaction) != 1 {
		return nil, err
	}

	return &results.Appqoeaction[0], nil
}

func (c *NitroClient) ListAppqoeaction() ([]AppqoeactionResource, error) {
	var results struct {
		Appqoeaction []AppqoeactionResource
	}

	if err := c.listResources("appqoeaction", &results); err != nil {
		return nil, err
	}

	return results.Appqoeaction, nil
}

func (c *NitroClient) AddAppqoeaction(resource AppqoeactionResource) error {
	return c.addResource("appqoeaction", resource)
}

func (c *NitroClient) RenameAppqoeaction(name string, newName string) error {
	return c.renameResource("appqoeaction", "name", name, newName)
}

func (c *NitroClient) UnsetAppqoeaction(name string, fields ...string) error {
	return c.unsetResource("appqoeaction", "name", name, fields)
}

func (c *NitroClient) UpdateAppqoeaction(resource AppqoeactionResource) error {
	update := appqoeaction_update{
		resource.Name,
		resource.Priority,
		resource.Altcontentsvcname,
		resource.Altcontentpath,
		resource.Polqdepth,
		resource.Priqdepth,
		resource.Maxconn,
		resource.Delay,
		resource.Dostrigexpression,
		resource.Dosaction,
		resource.Tcpprofile,
	}

	return c.Put("appqoeaction", update)
}
