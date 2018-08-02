package nitro

type ResponderpolicylabelResource struct {
	Labelname       string `json:"labelname"`
	Comment         string `json:"comment,omitempty"`
	Policylabeltype string `json:"policylabeltype,omitempty"`
}

type ResponderpolicylabelKey struct {
	Labelname string
}

type responderpolicylabel_update struct {
	Labelname string `json:"labelname"`
}

type responderpolicylabel_payload struct {
	responderpolicylabel interface{}
}

func responderpolicylabel_key_to_args(key ResponderpolicylabelKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteResponderpolicylabel(key ResponderpolicylabelKey) error {
	return c.deleteResourceWithArgs("responderpolicylabel", key.Labelname, responderpolicylabel_key_to_args(key))
}

func (c *NitroClient) GetResponderpolicylabel(key ResponderpolicylabelKey) (*ResponderpolicylabelResource, error) {
	var results struct {
		Responderpolicylabel []ResponderpolicylabelResource
	}

	if err := c.getResourceWithArgs("responderpolicylabel", key.Labelname, responderpolicylabel_key_to_args(key), &results); err != nil || len(results.Responderpolicylabel) != 1 {
		return nil, err
	}

	return &results.Responderpolicylabel[0], nil
}

func (c *NitroClient) ListResponderpolicylabel() ([]ResponderpolicylabelResource, error) {
	var results struct {
		Responderpolicylabel []ResponderpolicylabelResource
	}

	if err := c.listResources("responderpolicylabel", &results); err != nil {
		return nil, err
	}

	return results.Responderpolicylabel, nil
}

func (c *NitroClient) AddResponderpolicylabel(resource ResponderpolicylabelResource) error {
	return c.addResource("responderpolicylabel", resource)
}

func (c *NitroClient) RenameResponderpolicylabel(labelname string, newName string) error {
	return c.renameResource("responderpolicylabel", "labelname", labelname, newName)
}
