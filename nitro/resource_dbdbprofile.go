package nitro

type Dbdbprofile struct {
	Name                   string `json:"name"`
	Conmultiplex           string `json:"conmultiplex,omitempty"`
	Enablecachingconmuxoff string `json:"enablecachingconmuxoff,omitempty"`
	Interpretquery         string `json:"interpretquery,omitempty"`
	Kcdaccount             string `json:"kcdaccount,omitempty"`
	Stickiness             string `json:"stickiness,omitempty"`
}

type DbdbprofileKey struct {
	Name string
}

type dbdbprofile_update struct {
	Name                   string `json:"name"`
	Interpretquery         string `json:"interpretquery,omitempty"`
	Stickiness             string `json:"stickiness,omitempty"`
	Kcdaccount             string `json:"kcdaccount,omitempty"`
	Conmultiplex           string `json:"conmultiplex,omitempty"`
	Enablecachingconmuxoff string `json:"enablecachingconmuxoff,omitempty"`
}

type dbdbprofile_payload struct {
	dbdbprofile interface{}
}

func dbdbprofile_key_to_args(key DbdbprofileKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteDbdbprofile(key DbdbprofileKey) error {
	return c.deleteResourceWithArgs("dbdbprofile", key.Name, dbdbprofile_key_to_args(key))
}

func (c *NitroClient) GetDbdbprofile(key DbdbprofileKey) (*Dbdbprofile, error) {
	var results struct {
		Dbdbprofile []Dbdbprofile
	}

	if err := c.getResourceWithArgs("dbdbprofile", key.Name, dbdbprofile_key_to_args(key), &results); err != nil || len(results.Dbdbprofile) != 1 {
		return nil, err
	}

	return &results.Dbdbprofile[0], nil
}

func (c *NitroClient) ListDbdbprofile() ([]Dbdbprofile, error) {
	var results struct {
		Dbdbprofile []Dbdbprofile
	}

	if err := c.listResources("dbdbprofile", &results); err != nil {
		return nil, err
	}

	return results.Dbdbprofile, nil
}

func (c *NitroClient) AddDbdbprofile(resource Dbdbprofile) error {
	return c.addResource("dbdbprofile", resource)
}

func (c *NitroClient) RenameDbdbprofile(name string, newName string) error {
	return c.renameResource("dbdbprofile", "name", name, newName)
}

func (c *NitroClient) UnsetDbdbprofile(name string, fields ...string) error {
	return c.unsetResource("dbdbprofile", "name", name, fields)
}

func (c *NitroClient) UpdateDbdbprofile(resource Dbdbprofile) error {
	update := dbdbprofile_update{
		resource.Name,
		resource.Interpretquery,
		resource.Stickiness,
		resource.Kcdaccount,
		resource.Conmultiplex,
		resource.Enablecachingconmuxoff,
	}

	return c.Put("dbdbprofile", update)
}
