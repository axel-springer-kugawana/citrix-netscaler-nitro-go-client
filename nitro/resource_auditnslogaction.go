package nitro

type AuditnslogactionResource struct {
	Name                string   `json:"name"`
	Acl                 string   `json:"acl,omitempty"`
	Alg                 string   `json:"alg,omitempty"`
	Appflowexport       string   `json:"appflowexport,omitempty"`
	Dateformat          string   `json:"dateformat,omitempty"`
	Domainresolveretry  int      `json:"domainresolveretry,omitempty"`
	Logfacility         string   `json:"logfacility,omitempty"`
	Loglevel            []string `json:"loglevel,omitempty"`
	Lsn                 string   `json:"lsn,omitempty"`
	Serverdomainname    string   `json:"serverdomainname,omitempty"`
	Serverip            string   `json:"serverip,omitempty"`
	Serverport          int      `json:"serverport,omitempty"`
	Sslinterception     string   `json:"sslinterception,omitempty"`
	Subscriberlog       string   `json:"subscriberlog,omitempty"`
	Tcp                 string   `json:"tcp,omitempty"`
	Timezone            string   `json:"timezone,omitempty"`
	Userdefinedauditlog string   `json:"userdefinedauditlog,omitempty"`
}

type AuditnslogactionKey struct {
	Name string
}

type auditnslogaction_update struct {
	Name                string   `json:"name"`
	Serverip            string   `json:"serverip,omitempty"`
	Serverdomainname    string   `json:"serverdomainname,omitempty"`
	Domainresolveretry  int      `json:"domainresolveretry,omitempty"`
	Serverport          int      `json:"serverport,omitempty"`
	Loglevel            []string `json:"loglevel,omitempty"`
	Dateformat          string   `json:"dateformat,omitempty"`
	Logfacility         string   `json:"logfacility,omitempty"`
	Tcp                 string   `json:"tcp,omitempty"`
	Acl                 string   `json:"acl,omitempty"`
	Timezone            string   `json:"timezone,omitempty"`
	Userdefinedauditlog string   `json:"userdefinedauditlog,omitempty"`
	Appflowexport       string   `json:"appflowexport,omitempty"`
	Lsn                 string   `json:"lsn,omitempty"`
	Alg                 string   `json:"alg,omitempty"`
	Subscriberlog       string   `json:"subscriberlog,omitempty"`
	Sslinterception     string   `json:"sslinterception,omitempty"`
}

type auditnslogaction_payload struct {
	auditnslogaction interface{}
}

func auditnslogaction_key_to_args(key AuditnslogactionKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteAuditnslogaction(key AuditnslogactionKey) error {
	return c.deleteResourceWithArgs("auditnslogaction", key.Name, auditnslogaction_key_to_args(key))
}

func (c *NitroClient) GetAuditnslogaction(key AuditnslogactionKey) (*AuditnslogactionResource, error) {
	var results struct {
		Auditnslogaction []AuditnslogactionResource
	}

	if err := c.getResourceWithArgs("auditnslogaction", key.Name, auditnslogaction_key_to_args(key), &results); err != nil || len(results.Auditnslogaction) != 1 {
		return nil, err
	}

	return &results.Auditnslogaction[0], nil
}

func (c *NitroClient) ListAuditnslogaction() ([]AuditnslogactionResource, error) {
	var results struct {
		Auditnslogaction []AuditnslogactionResource
	}

	if err := c.listResources("auditnslogaction", &results); err != nil {
		return nil, err
	}

	return results.Auditnslogaction, nil
}

func (c *NitroClient) AddAuditnslogaction(resource AuditnslogactionResource) error {
	return c.addResource("auditnslogaction", resource)
}

func (c *NitroClient) RenameAuditnslogaction(name string, newName string) error {
	return c.renameResource("auditnslogaction", "name", name, newName)
}

func (c *NitroClient) UnsetAuditnslogaction(name string, fields ...string) error {
	return c.unsetResource("auditnslogaction", "name", name, fields)
}

func (c *NitroClient) UpdateAuditnslogaction(resource AuditnslogactionResource) error {
	update := auditnslogaction_update{
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
	}

	return c.Put("auditnslogaction", update)
}
