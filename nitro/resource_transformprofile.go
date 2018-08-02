package nitro

type TransformprofileResource struct {
	Name                      string `json:"name"`
	Comment                   string `json:"comment,omitempty"`
	Onlytransformabsurlinbody string `json:"onlytransformabsurlinbody,omitempty"`
	Type                      string `json:"type,omitempty"`
}

type TransformprofileKey struct {
	Name string
}

type transformprofile_update struct {
	Name                      string `json:"name"`
	Type                      string `json:"type,omitempty"`
	Onlytransformabsurlinbody string `json:"onlytransformabsurlinbody,omitempty"`
	Comment                   string `json:"comment,omitempty"`
}

type transformprofile_payload struct {
	transformprofile interface{}
}

func transformprofile_key_to_args(key TransformprofileKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteTransformprofile(key TransformprofileKey) error {
	return c.deleteResourceWithArgs("transformprofile", key.Name, transformprofile_key_to_args(key))
}

func (c *NitroClient) GetTransformprofile(key TransformprofileKey) (*TransformprofileResource, error) {
	var results struct {
		Transformprofile []TransformprofileResource
	}

	if err := c.getResourceWithArgs("transformprofile", key.Name, transformprofile_key_to_args(key), &results); err != nil || len(results.Transformprofile) != 1 {
		return nil, err
	}

	return &results.Transformprofile[0], nil
}

func (c *NitroClient) ListTransformprofile() ([]TransformprofileResource, error) {
	var results struct {
		Transformprofile []TransformprofileResource
	}

	if err := c.listResources("transformprofile", &results); err != nil {
		return nil, err
	}

	return results.Transformprofile, nil
}

func (c *NitroClient) AddTransformprofile(resource TransformprofileResource) error {
	return c.addResource("transformprofile", resource)
}

func (c *NitroClient) RenameTransformprofile(name string, newName string) error {
	return c.renameResource("transformprofile", "name", name, newName)
}

func (c *NitroClient) UnsetTransformprofile(name string, fields ...string) error {
	return c.unsetResource("transformprofile", "name", name, fields)
}

func (c *NitroClient) UpdateTransformprofile(resource TransformprofileResource) error {
	update := transformprofile_update{
		resource.Name,
		resource.Type,
		resource.Onlytransformabsurlinbody,
		resource.Comment,
	}

	return c.Put("transformprofile", update)
}
