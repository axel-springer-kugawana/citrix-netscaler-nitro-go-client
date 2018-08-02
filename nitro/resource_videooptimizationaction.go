package nitro

type Videooptimizationaction struct {
	Name    string `json:"name"`
	Comment string `json:"comment,omitempty"`
	Rate    int    `json:"rate,omitempty"`
	Type    string `json:"type,omitempty"`
}

type VideooptimizationactionKey struct {
	Name string
}

type videooptimizationaction_update struct {
	Name    string `json:"name"`
	Type    string `json:"type,omitempty"`
	Rate    int    `json:"rate,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type videooptimizationaction_payload struct {
	videooptimizationaction interface{}
}

func videooptimizationaction_key_to_args(key VideooptimizationactionKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteVideooptimizationaction(key VideooptimizationactionKey) error {
	return c.deleteResourceWithArgs("videooptimizationaction", key.Name, videooptimizationaction_key_to_args(key))
}

func (c *NitroClient) GetVideooptimizationaction(key VideooptimizationactionKey) (*Videooptimizationaction, error) {
	var results struct {
		Videooptimizationaction []Videooptimizationaction
	}

	if err := c.getResourceWithArgs("videooptimizationaction", key.Name, videooptimizationaction_key_to_args(key), &results); err != nil || len(results.Videooptimizationaction) != 1 {
		return nil, err
	}

	return &results.Videooptimizationaction[0], nil
}

func (c *NitroClient) ListVideooptimizationaction() ([]Videooptimizationaction, error) {
	var results struct {
		Videooptimizationaction []Videooptimizationaction
	}

	if err := c.listResources("videooptimizationaction", &results); err != nil {
		return nil, err
	}

	return results.Videooptimizationaction, nil
}

func (c *NitroClient) AddVideooptimizationaction(resource Videooptimizationaction) error {
	return c.addResource("videooptimizationaction", resource)
}

func (c *NitroClient) RenameVideooptimizationaction(name string, newName string) error {
	return c.renameResource("videooptimizationaction", "name", name, newName)
}

func (c *NitroClient) UnsetVideooptimizationaction(name string, fields ...string) error {
	return c.unsetResource("videooptimizationaction", "name", name, fields)
}

func (c *NitroClient) UpdateVideooptimizationaction(resource Videooptimizationaction) error {
	update := videooptimizationaction_update{
		resource.Name,
		resource.Type,
		resource.Rate,
		resource.Comment,
	}

	return c.Put("videooptimizationaction", update)
}
