package nitro

type AuthorizationpolicylabelResource struct {
	Labelname string `json:"labelname"`
}

type AuthorizationpolicylabelKey struct {
	Labelname string
}

type authorizationpolicylabel_update struct {
	Labelname string `json:"labelname"`
}

type authorizationpolicylabel_payload struct {
	authorizationpolicylabel interface{}
}

func authorizationpolicylabel_key_to_args(key AuthorizationpolicylabelKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteAuthorizationpolicylabel(key AuthorizationpolicylabelKey) error {
	return c.deleteResourceWithArgs("authorizationpolicylabel", key.Labelname, authorizationpolicylabel_key_to_args(key))
}

func (c *NitroClient) GetAuthorizationpolicylabel(key AuthorizationpolicylabelKey) (*AuthorizationpolicylabelResource, error) {
	var results struct {
		Authorizationpolicylabel []AuthorizationpolicylabelResource
	}

	if err := c.getResourceWithArgs("authorizationpolicylabel", key.Labelname, authorizationpolicylabel_key_to_args(key), &results); err != nil || len(results.Authorizationpolicylabel) != 1 {
		return nil, err
	}

	return &results.Authorizationpolicylabel[0], nil
}

func (c *NitroClient) ListAuthorizationpolicylabel() ([]AuthorizationpolicylabelResource, error) {
	var results struct {
		Authorizationpolicylabel []AuthorizationpolicylabelResource
	}

	if err := c.listResources("authorizationpolicylabel", &results); err != nil {
		return nil, err
	}

	return results.Authorizationpolicylabel, nil
}

func (c *NitroClient) AddAuthorizationpolicylabel(resource AuthorizationpolicylabelResource) error {
	return c.addResource("authorizationpolicylabel", resource)
}

func (c *NitroClient) RenameAuthorizationpolicylabel(labelname string, newName string) error {
	return c.renameResource("authorizationpolicylabel", "labelname", labelname, newName)
}
