package nitro

type AuditsyslogactionResource struct {
	Name                 string   `json:"name"`
	Acl                  string   `json:"acl,omitempty"`
	Alg                  string   `json:"alg,omitempty"`
	Appflowexport        string   `json:"appflowexport,omitempty"`
	Dateformat           string   `json:"dateformat,omitempty"`
	Dns                  string   `json:"dns,omitempty"`
	Domainresolveretry   int      `json:"domainresolveretry,omitempty"`
	Lbvservername        string   `json:"lbvservername,omitempty"`
	Logfacility          string   `json:"logfacility,omitempty"`
	Loglevel             []string `json:"loglevel,omitempty"`
	Lsn                  string   `json:"lsn,omitempty"`
	Maxlogdatasizetohold int      `json:"maxlogdatasizetohold,string,omitempty"`
	Netprofile           string   `json:"netprofile,omitempty"`
	Serverdomainname     string   `json:"serverdomainname,omitempty"`
	Serverip             string   `json:"serverip,omitempty"`
	Serverport           int      `json:"serverport,omitempty"`
	Sslinterception      string   `json:"sslinterception,omitempty"`
	Subscriberlog        string   `json:"subscriberlog,omitempty"`
	Tcp                  string   `json:"tcp,omitempty"`
	Tcpprofilename       string   `json:"tcpprofilename,omitempty"`
	Timezone             string   `json:"timezone,omitempty"`
	Transport            string   `json:"transport,omitempty"`
	Userdefinedauditlog  string   `json:"userdefinedauditlog,omitempty"`
}

type AuditsyslogactionKey struct {
	Name string
}

type auditsyslogaction_update struct {
	Name                 string   `json:"name"`
	Serverip             string   `json:"serverip,omitempty"`
	Serverdomainname     string   `json:"serverdomainname,omitempty"`
	Domainresolveretry   int      `json:"domainresolveretry,omitempty"`
	Serverport           int      `json:"serverport,omitempty"`
	Loglevel             []string `json:"loglevel,omitempty"`
	Dateformat           string   `json:"dateformat,omitempty"`
	Logfacility          string   `json:"logfacility,omitempty"`
	Tcp                  string   `json:"tcp,omitempty"`
	Acl                  string   `json:"acl,omitempty"`
	Timezone             string   `json:"timezone,omitempty"`
	Userdefinedauditlog  string   `json:"userdefinedauditlog,omitempty"`
	Appflowexport        string   `json:"appflowexport,omitempty"`
	Lsn                  string   `json:"lsn,omitempty"`
	Alg                  string   `json:"alg,omitempty"`
	Subscriberlog        string   `json:"subscriberlog,omitempty"`
	Sslinterception      string   `json:"sslinterception,omitempty"`
	Lbvservername        string   `json:"lbvservername,omitempty"`
	Tcpprofilename       string   `json:"tcpprofilename,omitempty"`
	Maxlogdatasizetohold int      `json:"maxlogdatasizetohold,string,omitempty"`
	Dns                  string   `json:"dns,omitempty"`
	Netprofile           string   `json:"netprofile,omitempty"`
}

type auditsyslogaction_payload struct {
	auditsyslogaction interface{}
}

func auditsyslogaction_key_to_args(key AuditsyslogactionKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteAuditsyslogaction(key AuditsyslogactionKey) error {
	return c.deleteResourceWithArgs("auditsyslogaction", key.Name, auditsyslogaction_key_to_args(key))
}

func (c *NitroClient) GetAuditsyslogaction(key AuditsyslogactionKey) (*AuditsyslogactionResource, error) {
	var results struct {
		Auditsyslogaction []AuditsyslogactionResource
	}

	if err := c.getResourceWithArgs("auditsyslogaction", key.Name, auditsyslogaction_key_to_args(key), &results); err != nil || len(results.Auditsyslogaction) != 1 {
		return nil, err
	}

	return &results.Auditsyslogaction[0], nil
}

func (c *NitroClient) ListAuditsyslogaction() ([]AuditsyslogactionResource, error) {
	var results struct {
		Auditsyslogaction []AuditsyslogactionResource
	}

	if err := c.listResources("auditsyslogaction", &results); err != nil {
		return nil, err
	}

	return results.Auditsyslogaction, nil
}

func (c *NitroClient) AddAuditsyslogaction(resource AuditsyslogactionResource) error {
	return c.addResource("auditsyslogaction", resource)
}

func (c *NitroClient) RenameAuditsyslogaction(name string, newName string) error {
	return c.renameResource("auditsyslogaction", "name", name, newName)
}

func (c *NitroClient) UnsetAuditsyslogaction(name string, fields ...string) error {
	return c.unsetResource("auditsyslogaction", "name", name, fields)
}

func (c *NitroClient) UpdateAuditsyslogaction(resource AuditsyslogactionResource) error {
	update := auditsyslogaction_update{
		resource.Name,
		resource.Serverip,
		resource.Serverdomainname,
		resource.Domainresolveretry,
		resource.Serverport,
		resource.Loglevel,
		resource.Dateformat,
		resource.Logfacility,
		resource.Tcp,
		resource.Acl,
		resource.Timezone,
		resource.Userdefinedauditlog,
		resource.Appflowexport,
		resource.Lsn,
		resource.Alg,
		resource.Subscriberlog,
		resource.Sslinterception,
		resource.Lbvservername,
		resource.Tcpprofilename,
		resource.Maxlogdatasizetohold,
		resource.Dns,
		resource.Netprofile,
	}

	return c.Put("auditsyslogaction", update)
}
