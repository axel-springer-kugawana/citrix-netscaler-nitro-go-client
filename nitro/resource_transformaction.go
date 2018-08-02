package nitro

type Transformaction struct {
	Name             string `json:"name"`
	Comment          string `json:"comment,omitempty"`
	Cookiedomainfrom string `json:"cookiedomainfrom,omitempty"`
	Cookiedomaininto string `json:"cookiedomaininto,omitempty"`
	Priority         int    `json:"priority,string,omitempty"`
	Profilename      string `json:"profilename,omitempty"`
	Requrlfrom       string `json:"requrlfrom,omitempty"`
	Requrlinto       string `json:"requrlinto,omitempty"`
	Resurlfrom       string `json:"resurlfrom,omitempty"`
	Resurlinto       string `json:"resurlinto,omitempty"`
	State            string `json:"state,omitempty"`
}

type TransformactionKey struct {
	Name string
}

type transformaction_update struct {
	Name             string `json:"name"`
	Priority         int    `json:"priority,string,omitempty"`
	Requrlfrom       string `json:"requrlfrom,omitempty"`
	Requrlinto       string `json:"requrlinto,omitempty"`
	Resurlfrom       string `json:"resurlfrom,omitempty"`
	Resurlinto       string `json:"resurlinto,omitempty"`
	Cookiedomainfrom string `json:"cookiedomainfrom,omitempty"`
	Cookiedomaininto string `json:"cookiedomaininto,omitempty"`
	State            string `json:"state,omitempty"`
	Comment          string `json:"comment,omitempty"`
}

type transformaction_payload struct {
	transformaction interface{}
}

func transformaction_key_to_args(key TransformactionKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteTransformaction(key TransformactionKey) error {
	return c.deleteResourceWithArgs("transformaction", key.Name, transformaction_key_to_args(key))
}

func (c *NitroClient) GetTransformaction(key TransformactionKey) (*Transformaction, error) {
	var results struct {
		Transformaction []Transformaction
	}

	if err := c.getResourceWithArgs("transformaction", key.Name, transformaction_key_to_args(key), &results); err != nil || len(results.Transformaction) != 1 {
		return nil, err
	}

	return &results.Transformaction[0], nil
}

func (c *NitroClient) ListTransformaction() ([]Transformaction, error) {
	var results struct {
		Transformaction []Transformaction
	}

	if err := c.listResources("transformaction", &results); err != nil {
		return nil, err
	}

	return results.Transformaction, nil
}

func (c *NitroClient) AddTransformaction(resource Transformaction) error {
	return c.addResource("transformaction", resource)
}

func (c *NitroClient) RenameTransformaction(name string, newName string) error {
	return c.renameResource("transformaction", "name", name, newName)
}

func (c *NitroClient) UnsetTransformaction(name string, fields ...string) error {
	return c.unsetResource("transformaction", "name", name, fields)
}

func (c *NitroClient) UpdateTransformaction(resource Transformaction) error {
	update := transformaction_update{
		resource.Name,
		resource.Priority,
		resource.Requrlfrom,
		resource.Requrlinto,
		resource.Resurlfrom,
		resource.Resurlinto,
		resource.Cookiedomainfrom,
		resource.Cookiedomaininto,
		resource.State,
		resource.Comment,
	}

	return c.Put("transformaction", update)
}
