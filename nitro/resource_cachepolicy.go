package nitro

type CachepolicyResource struct {
	Policyname   string   `json:"policyname"`
	Action       string   `json:"action,omitempty"`
	Invalgroups  []string `json:"invalgroups,omitempty"`
	Invalobjects []string `json:"invalobjects,omitempty"`
	Rule         string   `json:"rule,omitempty"`
	Storeingroup string   `json:"storeingroup,omitempty"`
	Undefaction  string   `json:"undefaction,omitempty"`
}

type CachepolicyKey struct {
	Policyname string
}

type cachepolicy_update struct {
	Policyname   string   `json:"policyname"`
	Rule         string   `json:"rule,omitempty"`
	Action       string   `json:"action,omitempty"`
	Storeingroup string   `json:"storeingroup,omitempty"`
	Invalgroups  []string `json:"invalgroups,omitempty"`
	Invalobjects []string `json:"invalobjects,omitempty"`
	Undefaction  string   `json:"undefaction,omitempty"`
}

type cachepolicy_payload struct {
	cachepolicy interface{}
}

func cachepolicy_key_to_args(key CachepolicyKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteCachepolicy(key CachepolicyKey) error {
	return c.deleteResourceWithArgs("cachepolicy", key.Policyname, cachepolicy_key_to_args(key))
}

func (c *NitroClient) GetCachepolicy(key CachepolicyKey) (*CachepolicyResource, error) {
	var results struct {
		Cachepolicy []CachepolicyResource
	}

	if err := c.getResourceWithArgs("cachepolicy", key.Policyname, cachepolicy_key_to_args(key), &results); err != nil || len(results.Cachepolicy) != 1 {
		return nil, err
	}

	return &results.Cachepolicy[0], nil
}

func (c *NitroClient) ListCachepolicy() ([]CachepolicyResource, error) {
	var results struct {
		Cachepolicy []CachepolicyResource
	}

	if err := c.listResources("cachepolicy", &results); err != nil {
		return nil, err
	}

	return results.Cachepolicy, nil
}

func (c *NitroClient) AddCachepolicy(resource CachepolicyResource) error {
	return c.addResource("cachepolicy", resource)
}

func (c *NitroClient) RenameCachepolicy(policyname string, newName string) error {
	return c.renameResource("cachepolicy", "policyname", policyname, newName)
}

func (c *NitroClient) UnsetCachepolicy(policyname string, fields ...string) error {
	return c.unsetResource("cachepolicy", "policyname", policyname, fields)
}

func (c *NitroClient) UpdateCachepolicy(resource CachepolicyResource) error {
	update := cachepolicy_update{
		resource.Policyname,
		resource.Rule,
		resource.Action,
		resource.Storeingroup,
		resource.Invalgroups,
		resource.Invalobjects,
		resource.Undefaction,
	}

	return c.Put("cachepolicy", update)
}
