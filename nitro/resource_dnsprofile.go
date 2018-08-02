package nitro

type DnsprofileResource struct {
	Dnsprofilename         string `json:"dnsprofilename"`
	Cacheecsresponses      string `json:"cacheecsresponses,omitempty"`
	Cachenegativeresponses string `json:"cachenegativeresponses,omitempty"`
	Cacherecords           string `json:"cacherecords,omitempty"`
	Dnsanswerseclogging    string `json:"dnsanswerseclogging,omitempty"`
	Dnserrorlogging        string `json:"dnserrorlogging,omitempty"`
	Dnsextendedlogging     string `json:"dnsextendedlogging,omitempty"`
	Dnsquerylogging        string `json:"dnsquerylogging,omitempty"`
	Dropmultiqueryrequest  string `json:"dropmultiqueryrequest,omitempty"`
}

type DnsprofileKey struct {
	Dnsprofilename string
}

type dnsprofile_update struct {
	Dnsprofilename         string `json:"dnsprofilename"`
	Dnsquerylogging        string `json:"dnsquerylogging,omitempty"`
	Dnsanswerseclogging    string `json:"dnsanswerseclogging,omitempty"`
	Dnsextendedlogging     string `json:"dnsextendedlogging,omitempty"`
	Dnserrorlogging        string `json:"dnserrorlogging,omitempty"`
	Cacherecords           string `json:"cacherecords,omitempty"`
	Cachenegativeresponses string `json:"cachenegativeresponses,omitempty"`
	Dropmultiqueryrequest  string `json:"dropmultiqueryrequest,omitempty"`
	Cacheecsresponses      string `json:"cacheecsresponses,omitempty"`
}

type dnsprofile_payload struct {
	dnsprofile interface{}
}

func dnsprofile_key_to_args(key DnsprofileKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteDnsprofile(key DnsprofileKey) error {
	return c.deleteResourceWithArgs("dnsprofile", key.Dnsprofilename, dnsprofile_key_to_args(key))
}

func (c *NitroClient) GetDnsprofile(key DnsprofileKey) (*DnsprofileResource, error) {
	var results struct {
		Dnsprofile []DnsprofileResource
	}

	if err := c.getResourceWithArgs("dnsprofile", key.Dnsprofilename, dnsprofile_key_to_args(key), &results); err != nil || len(results.Dnsprofile) != 1 {
		return nil, err
	}

	return &results.Dnsprofile[0], nil
}

func (c *NitroClient) ListDnsprofile() ([]DnsprofileResource, error) {
	var results struct {
		Dnsprofile []DnsprofileResource
	}

	if err := c.listResources("dnsprofile", &results); err != nil {
		return nil, err
	}

	return results.Dnsprofile, nil
}

func (c *NitroClient) AddDnsprofile(resource DnsprofileResource) error {
	return c.addResource("dnsprofile", resource)
}

func (c *NitroClient) RenameDnsprofile(dnsprofilename string, newName string) error {
	return c.renameResource("dnsprofile", "dnsprofilename", dnsprofilename, newName)
}

func (c *NitroClient) UnsetDnsprofile(dnsprofilename string, fields ...string) error {
	return c.unsetResource("dnsprofile", "dnsprofilename", dnsprofilename, fields)
}

func (c *NitroClient) UpdateDnsprofile(resource DnsprofileResource) error {
	update := dnsprofile_update{
		resource.Dnsprofilename,
		resource.Dnsquerylogging,
		resource.Dnsanswerseclogging,
		resource.Dnsextendedlogging,
		resource.Dnserrorlogging,
		resource.Cacherecords,
		resource.Cachenegativeresponses,
		resource.Dropmultiqueryrequest,
		resource.Cacheecsresponses,
	}

	return c.Put("dnsprofile", update)
}
