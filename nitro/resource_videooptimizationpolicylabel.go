package nitro

type VideooptimizationpolicylabelResource struct {
	Labelname       string `json:"labelname"`
	Comment         string `json:"comment,omitempty"`
	Policylabeltype string `json:"policylabeltype,omitempty"`
}

type VideooptimizationpolicylabelKey struct {
	Labelname string
}

type videooptimizationpolicylabel_update struct {
	Labelname string `json:"labelname"`
}

type videooptimizationpolicylabel_payload struct {
	videooptimizationpolicylabel interface{}
}

func videooptimizationpolicylabel_key_to_args(key VideooptimizationpolicylabelKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteVideooptimizationpolicylabel(key VideooptimizationpolicylabelKey) error {
	return c.deleteResourceWithArgs("videooptimizationpolicylabel", key.Labelname, videooptimizationpolicylabel_key_to_args(key))
}

func (c *NitroClient) GetVideooptimizationpolicylabel(key VideooptimizationpolicylabelKey) (*VideooptimizationpolicylabelResource, error) {
	var results struct {
		Videooptimizationpolicylabel []VideooptimizationpolicylabelResource
	}

	if err := c.getResourceWithArgs("videooptimizationpolicylabel", key.Labelname, videooptimizationpolicylabel_key_to_args(key), &results); err != nil || len(results.Videooptimizationpolicylabel) != 1 {
		return nil, err
	}

	return &results.Videooptimizationpolicylabel[0], nil
}

func (c *NitroClient) ListVideooptimizationpolicylabel() ([]VideooptimizationpolicylabelResource, error) {
	var results struct {
		Videooptimizationpolicylabel []VideooptimizationpolicylabelResource
	}

	if err := c.listResources("videooptimizationpolicylabel", &results); err != nil {
		return nil, err
	}

	return results.Videooptimizationpolicylabel, nil
}

func (c *NitroClient) AddVideooptimizationpolicylabel(resource VideooptimizationpolicylabelResource) error {
	return c.addResource("videooptimizationpolicylabel", resource)
}

func (c *NitroClient) RenameVideooptimizationpolicylabel(labelname string, newName string) error {
	return c.renameResource("videooptimizationpolicylabel", "labelname", labelname, newName)
}
