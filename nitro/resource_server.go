package nitro

type Server struct {
	Name               string `json:"name"`
	State              string `json:"state,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Domain             string `json:"domain,omitempty"`
	Domainresolveretry int    `json:"domainresolveretry,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Ipv6address        string `json:"ipv6address,omitempty"`
	Td                 int    `json:"td,string,omitempty"`
	Translationip      string `json:"translationip,omitempty"`
	Translationmask    string `json:"translationmask,omitempty"`
}

type ServerKey struct {
	Name string
}

type server_update struct {
	Name               string `json:"name"`
	Comment            string `json:"comment,omitempty"`
	Domainresolveretry int    `json:"domainresolveretry,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Translationip      string `json:"translationip,omitempty"`
	Translationmask    string `json:"translationmask,omitempty"`
}

type server_payload struct {
	server interface{}
}

func server_key_to_args(key ServerKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteServer(key ServerKey) error {
	return c.deleteResourceWithArgs("server", key.Name, server_key_to_args(key))
}

func (c *NitroClient) GetServer(key ServerKey) (*Server, error) {
	var results struct {
		Server []Server
	}

	if err := c.getResourceWithArgs("server", key.Name, server_key_to_args(key), &results); err != nil || len(results.Server) != 1 {
		return nil, err
	}

	return &results.Server[0], nil
}

func (c *NitroClient) ListServer() ([]Server, error) {
	var results struct {
		Server []Server
	}

	if err := c.listResources("server", &results); err != nil {
		return nil, err
	}

	return results.Server, nil
}

func (c *NitroClient) AddServer(resource Server) error {
	return c.addResource("server", resource)
}

func (c *NitroClient) RenameServer(name string, newName string) error {
	return c.renameResource("server", "name", name, newName)
}

func (c *NitroClient) UnsetServer(name string, fields ...string) error {
	return c.unsetResource("server", "name", name, fields)
}

func (c *NitroClient) UpdateServer(resource Server) error {
	update := server_update{
		resource.Name,
		resource.Comment,
		resource.Domainresolveretry,
		resource.Ipaddress,
		resource.Translationip,
		resource.Translationmask,
	}

	return c.Put("server", update)
}
func (c *NitroClient) EnableServer(name string) error {
	return c.enableResource("server", "name", name)
}

func (c *NitroClient) DisableServer(name string) error {
	return c.disableResource("server", "name", name)
}
