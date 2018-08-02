package nitro

type Videooptimizationpolicy struct {
	Name        string `json:"name"`
	Action      string `json:"action,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
}

type VideooptimizationpolicyKey struct {
	Name string
}

type videooptimizationpolicy_update struct {
	Name        string `json:"name"`
	Rule        string `json:"rule,omitempty"`
	Action      string `json:"action,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
}

type videooptimizationpolicy_payload struct {
	videooptimizationpolicy interface{}
}

func videooptimizationpolicy_key_to_args(key VideooptimizationpolicyKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteVideooptimizationpolicy(key VideooptimizationpolicyKey) error {
	return c.deleteResourceWithArgs("videooptimizationpolicy", key.Name, videooptimizationpolicy_key_to_args(key))
}

func (c *NitroClient) GetVideooptimizationpolicy(key VideooptimizationpolicyKey) (*Videooptimizationpolicy, error) {
	var results struct {
		Videooptimizationpolicy []Videooptimizationpolicy
	}

	if err := c.getResourceWithArgs("videooptimizationpolicy", key.Name, videooptimizationpolicy_key_to_args(key), &results); err != nil || len(results.Videooptimizationpolicy) != 1 {
		return nil, err
	}

	return &results.Videooptimizationpolicy[0], nil
}

func (c *NitroClient) ListVideooptimizationpolicy() ([]Videooptimizationpolicy, error) {
	var results struct {
		Videooptimizationpolicy []Videooptimizationpolicy
	}

	if err := c.listResources("videooptimizationpolicy", &results); err != nil {
		return nil, err
	}

	return results.Videooptimizationpolicy, nil
}

func (c *NitroClient) AddVideooptimizationpolicy(resource Videooptimizationpolicy) error {
	return c.addResource("videooptimizationpolicy", resource)
}

func (c *NitroClient) RenameVideooptimizationpolicy(name string, newName string) error {
	return c.renameResource("videooptimizationpolicy", "name", name, newName)
}

func (c *NitroClient) UnsetVideooptimizationpolicy(name string, fields ...string) error {
	return c.unsetResource("videooptimizationpolicy", "name", name, fields)
}

func (c *NitroClient) UpdateVideooptimizationpolicy(resource Videooptimizationpolicy) error {
	update := videooptimizationpolicy_update{
		resource.Name,
		resource.Rule,
		resource.Action,
		resource.Undefaction,
		resource.Comment,
		resource.Logaction,
	}

	return c.Put("videooptimizationpolicy", update)
}
