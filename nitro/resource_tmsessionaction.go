package nitro

type TmsessionactionResource struct {
	Name                       string `json:"name"`
	Defaultauthorizationaction string `json:"defaultauthorizationaction,omitempty"`
	Homepage                   string `json:"homepage,omitempty"`
	Httponlycookie             string `json:"httponlycookie,omitempty"`
	Kcdaccount                 string `json:"kcdaccount,omitempty"`
	Persistentcookie           string `json:"persistentcookie,omitempty"`
	Persistentcookievalidity   int    `json:"persistentcookievalidity,string,omitempty"`
	Sesstimeout                int    `json:"sesstimeout,string,omitempty"`
	Sso                        string `json:"sso,omitempty"`
	Ssocredential              string `json:"ssocredential,omitempty"`
	Ssodomain                  string `json:"ssodomain,omitempty"`
}

type TmsessionactionKey struct {
	Name string
}

type tmsessionaction_update struct {
	Name                       string `json:"name"`
	Sesstimeout                int    `json:"sesstimeout,string,omitempty"`
	Defaultauthorizationaction string `json:"defaultauthorizationaction,omitempty"`
	Sso                        string `json:"sso,omitempty"`
	Ssocredential              string `json:"ssocredential,omitempty"`
	Ssodomain                  string `json:"ssodomain,omitempty"`
	Kcdaccount                 string `json:"kcdaccount,omitempty"`
	Httponlycookie             string `json:"httponlycookie,omitempty"`
	Persistentcookie           string `json:"persistentcookie,omitempty"`
	Persistentcookievalidity   int    `json:"persistentcookievalidity,string,omitempty"`
	Homepage                   string `json:"homepage,omitempty"`
}

type tmsessionaction_payload struct {
	tmsessionaction interface{}
}

func tmsessionaction_key_to_args(key TmsessionactionKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteTmsessionaction(key TmsessionactionKey) error {
	return c.deleteResourceWithArgs("tmsessionaction", key.Name, tmsessionaction_key_to_args(key))
}

func (c *NitroClient) GetTmsessionaction(key TmsessionactionKey) (*TmsessionactionResource, error) {
	var results struct {
		Tmsessionaction []TmsessionactionResource
	}

	if err := c.getResourceWithArgs("tmsessionaction", key.Name, tmsessionaction_key_to_args(key), &results); err != nil || len(results.Tmsessionaction) != 1 {
		return nil, err
	}

	return &results.Tmsessionaction[0], nil
}

func (c *NitroClient) ListTmsessionaction() ([]TmsessionactionResource, error) {
	var results struct {
		Tmsessionaction []TmsessionactionResource
	}

	if err := c.listResources("tmsessionaction", &results); err != nil {
		return nil, err
	}

	return results.Tmsessionaction, nil
}

func (c *NitroClient) AddTmsessionaction(resource TmsessionactionResource) error {
	return c.addResource("tmsessionaction", resource)
}

func (c *NitroClient) RenameTmsessionaction(name string, newName string) error {
	return c.renameResource("tmsessionaction", "name", name, newName)
}

func (c *NitroClient) UnsetTmsessionaction(name string, fields ...string) error {
	return c.unsetResource("tmsessionaction", "name", name, fields)
}

func (c *NitroClient) UpdateTmsessionaction(resource TmsessionactionResource) error {
	update := tmsessionaction_update{
		resource.Name,
		resource.Sesstimeout,
		resource.Defaultauthorizationaction,
		resource.Sso,
		resource.Ssocredential,
		resource.Ssodomain,
		resource.Kcdaccount,
		resource.Httponlycookie,
		resource.Persistentcookie,
		resource.Persistentcookievalidity,
		resource.Homepage,
	}

	return c.Put("tmsessionaction", update)
}
