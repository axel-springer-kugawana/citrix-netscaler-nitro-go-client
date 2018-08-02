package nitro

type LbprofileResource struct {
	Lbprofilename                 string `json:"lbprofilename"`
	Cookiepassphrase              string `json:"cookiepassphrase,omitempty"`
	Dbslb                         string `json:"dbslb,omitempty"`
	Httponlycookieflag            string `json:"httponlycookieflag,omitempty"`
	Processlocal                  string `json:"processlocal,omitempty"`
	Useencryptedpersistencecookie string `json:"useencryptedpersistencecookie,omitempty"`
	Usesecuredpersistencecookie   string `json:"usesecuredpersistencecookie,omitempty"`
}

type LbprofileKey struct {
	Lbprofilename string
}

type lbprofile_update struct {
	Lbprofilename                 string `json:"lbprofilename"`
	Cookiepassphrase              string `json:"cookiepassphrase,omitempty"`
	Dbslb                         string `json:"dbslb,omitempty"`
	Processlocal                  string `json:"processlocal,omitempty"`
	Httponlycookieflag            string `json:"httponlycookieflag,omitempty"`
	Usesecuredpersistencecookie   string `json:"usesecuredpersistencecookie,omitempty"`
	Useencryptedpersistencecookie string `json:"useencryptedpersistencecookie,omitempty"`
}

type lbprofile_payload struct {
	lbprofile interface{}
}

func lbprofile_key_to_args(key LbprofileKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteLbprofile(key LbprofileKey) error {
	return c.deleteResourceWithArgs("lbprofile", key.Lbprofilename, lbprofile_key_to_args(key))
}

func (c *NitroClient) GetLbprofile(key LbprofileKey) (*LbprofileResource, error) {
	var results struct {
		Lbprofile []LbprofileResource
	}

	if err := c.getResourceWithArgs("lbprofile", key.Lbprofilename, lbprofile_key_to_args(key), &results); err != nil || len(results.Lbprofile) != 1 {
		return nil, err
	}

	return &results.Lbprofile[0], nil
}

func (c *NitroClient) ListLbprofile() ([]LbprofileResource, error) {
	var results struct {
		Lbprofile []LbprofileResource
	}

	if err := c.listResources("lbprofile", &results); err != nil {
		return nil, err
	}

	return results.Lbprofile, nil
}

func (c *NitroClient) AddLbprofile(resource LbprofileResource) error {
	return c.addResource("lbprofile", resource)
}

func (c *NitroClient) RenameLbprofile(lbprofilename string, newName string) error {
	return c.renameResource("lbprofile", "lbprofilename", lbprofilename, newName)
}

func (c *NitroClient) UnsetLbprofile(lbprofilename string, fields ...string) error {
	return c.unsetResource("lbprofile", "lbprofilename", lbprofilename, fields)
}

func (c *NitroClient) UpdateLbprofile(resource LbprofileResource) error {
	update := lbprofile_update{
		resource.Lbprofilename,
		resource.Cookiepassphrase,
		resource.Dbslb,
		resource.Processlocal,
		resource.Httponlycookieflag,
		resource.Usesecuredpersistencecookie,
		resource.Useencryptedpersistencecookie,
	}

	return c.Put("lbprofile", update)
}
