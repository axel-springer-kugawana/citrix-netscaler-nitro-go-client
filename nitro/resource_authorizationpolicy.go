package nitro

type Authorizationpolicy struct {
	Name   string `json:"name"`
	Action string `json:"action,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

type AuthorizationpolicyKey struct {
	Name string
}

type authorizationpolicy_update struct {
	Name   string `json:"name"`
	Rule   string `json:"rule,omitempty"`
	Action string `json:"action,omitempty"`
}

type authorizationpolicy_payload struct {
	authorizationpolicy interface{}
}

func authorizationpolicy_key_to_args(key AuthorizationpolicyKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteAuthorizationpolicy(key AuthorizationpolicyKey) error {
	return c.deleteResourceWithArgs("authorizationpolicy", key.Name, authorizationpolicy_key_to_args(key))
}

func (c *NitroClient) GetAuthorizationpolicy(key AuthorizationpolicyKey) (*Authorizationpolicy, error) {
	var results struct {
		Authorizationpolicy []Authorizationpolicy
	}

	if err := c.getResourceWithArgs("authorizationpolicy", key.Name, authorizationpolicy_key_to_args(key), &results); err != nil || len(results.Authorizationpolicy) != 1 {
		return nil, err
	}

	return &results.Authorizationpolicy[0], nil
}

func (c *NitroClient) ListAuthorizationpolicy() ([]Authorizationpolicy, error) {
	var results struct {
		Authorizationpolicy []Authorizationpolicy
	}

	if err := c.listResources("authorizationpolicy", &results); err != nil {
		return nil, err
	}

	return results.Authorizationpolicy, nil
}

func (c *NitroClient) AddAuthorizationpolicy(resource Authorizationpolicy) error {
	return c.addResource("authorizationpolicy", resource)
}

func (c *NitroClient) RenameAuthorizationpolicy(name string, newName string) error {
	return c.renameResource("authorizationpolicy", "name", name, newName)
}

func (c *NitroClient) UnsetAuthorizationpolicy(name string, fields ...string) error {
	return c.unsetResource("authorizationpolicy", "name", name, fields)
}

func (c *NitroClient) UpdateAuthorizationpolicy(resource Authorizationpolicy) error {
	update := authorizationpolicy_update{
		resource.Name,
		resource.Rule,
		resource.Action,
	}

	return c.Put("authorizationpolicy", update)
}
