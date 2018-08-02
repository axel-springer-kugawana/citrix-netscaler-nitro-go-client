package nitro

type Lbgroup struct {
	Name                     string `json:"name"`
	Backuppersistencetimeout int    `json:"backuppersistencetimeout,string,omitempty"`
	Cookiedomain             string `json:"cookiedomain,omitempty"`
	Cookiename               string `json:"cookiename,omitempty"`
	Persistencebackup        string `json:"persistencebackup,omitempty"`
	Persistencetype          string `json:"persistencetype,omitempty"`
	Persistmask              string `json:"persistmask,omitempty"`
	Rule                     string `json:"rule,omitempty"`
	Timeout                  int    `json:"timeout,string,omitempty"`
	Usevserverpersistency    string `json:"usevserverpersistency,omitempty"`
	V6persistmasklen         int    `json:"v6persistmasklen,string,omitempty"`
}

type LbgroupKey struct {
	Name string
}

type lbgroup_update struct {
	Name                     string `json:"name"`
	Persistencetype          string `json:"persistencetype,omitempty"`
	Persistencebackup        string `json:"persistencebackup,omitempty"`
	Backuppersistencetimeout int    `json:"backuppersistencetimeout,string,omitempty"`
	Persistmask              string `json:"persistmask,omitempty"`
	Cookiename               string `json:"cookiename,omitempty"`
	V6persistmasklen         int    `json:"v6persistmasklen,string,omitempty"`
	Cookiedomain             string `json:"cookiedomain,omitempty"`
	Timeout                  int    `json:"timeout,string,omitempty"`
	Rule                     string `json:"rule,omitempty"`
	Usevserverpersistency    string `json:"usevserverpersistency,omitempty"`
}

type lbgroup_payload struct {
	lbgroup interface{}
}

func lbgroup_key_to_args(key LbgroupKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteLbgroup(key LbgroupKey) error {
	return c.deleteResourceWithArgs("lbgroup", key.Name, lbgroup_key_to_args(key))
}

func (c *NitroClient) GetLbgroup(key LbgroupKey) (*Lbgroup, error) {
	var results struct {
		Lbgroup []Lbgroup
	}

	if err := c.getResourceWithArgs("lbgroup", key.Name, lbgroup_key_to_args(key), &results); err != nil || len(results.Lbgroup) != 1 {
		return nil, err
	}

	return &results.Lbgroup[0], nil
}

func (c *NitroClient) ListLbgroup() ([]Lbgroup, error) {
	var results struct {
		Lbgroup []Lbgroup
	}

	if err := c.listResources("lbgroup", &results); err != nil {
		return nil, err
	}

	return results.Lbgroup, nil
}

func (c *NitroClient) AddLbgroup(resource Lbgroup) error {
	return c.addResource("lbgroup", resource)
}

func (c *NitroClient) RenameLbgroup(name string, newName string) error {
	return c.renameResource("lbgroup", "name", name, newName)
}

func (c *NitroClient) UnsetLbgroup(name string, fields ...string) error {
	return c.unsetResource("lbgroup", "name", name, fields)
}

func (c *NitroClient) UpdateLbgroup(resource Lbgroup) error {
	update := lbgroup_update{
		resource.Name,
		resource.Persistencetype,
		resource.Persistencebackup,
		resource.Backuppersistencetimeout,
		resource.Persistmask,
		resource.Cookiename,
		resource.V6persistmasklen,
		resource.Cookiedomain,
		resource.Timeout,
		resource.Rule,
		resource.Usevserverpersistency,
	}

	return c.Put("lbgroup", update)
}
