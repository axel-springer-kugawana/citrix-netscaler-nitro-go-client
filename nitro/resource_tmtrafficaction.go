package nitro

type TmtrafficactionResource struct {
	Name             string `json:"name"`
	Apptimeout       int    `json:"apptimeout,string,omitempty"`
	Forcedtimeout    string `json:"forcedtimeout,omitempty"`
	Forcedtimeoutval int    `json:"forcedtimeoutval,string,omitempty"`
	Formssoaction    string `json:"formssoaction,omitempty"`
	Initiatelogout   string `json:"initiatelogout,omitempty"`
	Kcdaccount       string `json:"kcdaccount,omitempty"`
	Passwdexpression string `json:"passwdexpression,omitempty"`
	Persistentcookie string `json:"persistentcookie,omitempty"`
	Samlssoprofile   string `json:"samlssoprofile,omitempty"`
	Sso              string `json:"sso,omitempty"`
	Userexpression   string `json:"userexpression,omitempty"`
}

type TmtrafficactionKey struct {
	Name string
}

type tmtrafficaction_update struct {
	Name             string `json:"name"`
	Apptimeout       int    `json:"apptimeout,string,omitempty"`
	Sso              string `json:"sso,omitempty"`
	Formssoaction    string `json:"formssoaction,omitempty"`
	Persistentcookie string `json:"persistentcookie,omitempty"`
	Initiatelogout   string `json:"initiatelogout,omitempty"`
	Kcdaccount       string `json:"kcdaccount,omitempty"`
	Samlssoprofile   string `json:"samlssoprofile,omitempty"`
	Forcedtimeout    string `json:"forcedtimeout,omitempty"`
	Forcedtimeoutval int    `json:"forcedtimeoutval,string,omitempty"`
	Userexpression   string `json:"userexpression,omitempty"`
	Passwdexpression string `json:"passwdexpression,omitempty"`
}

type tmtrafficaction_payload struct {
	tmtrafficaction interface{}
}

func tmtrafficaction_key_to_args(key TmtrafficactionKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteTmtrafficaction(key TmtrafficactionKey) error {
	return c.deleteResourceWithArgs("tmtrafficaction", key.Name, tmtrafficaction_key_to_args(key))
}

func (c *NitroClient) GetTmtrafficaction(key TmtrafficactionKey) (*TmtrafficactionResource, error) {
	var results struct {
		Tmtrafficaction []TmtrafficactionResource
	}

	if err := c.getResourceWithArgs("tmtrafficaction", key.Name, tmtrafficaction_key_to_args(key), &results); err != nil || len(results.Tmtrafficaction) != 1 {
		return nil, err
	}

	return &results.Tmtrafficaction[0], nil
}

func (c *NitroClient) ListTmtrafficaction() ([]TmtrafficactionResource, error) {
	var results struct {
		Tmtrafficaction []TmtrafficactionResource
	}

	if err := c.listResources("tmtrafficaction", &results); err != nil {
		return nil, err
	}

	return results.Tmtrafficaction, nil
}

func (c *NitroClient) AddTmtrafficaction(resource TmtrafficactionResource) error {
	return c.addResource("tmtrafficaction", resource)
}

func (c *NitroClient) RenameTmtrafficaction(name string, newName string) error {
	return c.renameResource("tmtrafficaction", "name", name, newName)
}

func (c *NitroClient) UnsetTmtrafficaction(name string, fields ...string) error {
	return c.unsetResource("tmtrafficaction", "name", name, fields)
}

func (c *NitroClient) UpdateTmtrafficaction(resource TmtrafficactionResource) error {
	update := tmtrafficaction_update{
		resource.Name,
		resource.Apptimeout,
		resource.Sso,
		resource.Formssoaction,
		resource.Persistentcookie,
		resource.Initiatelogout,
		resource.Kcdaccount,
		resource.Samlssoprofile,
		resource.Forcedtimeout,
		resource.Forcedtimeoutval,
		resource.Userexpression,
		resource.Passwdexpression,
	}

	return c.Put("tmtrafficaction", update)
}
