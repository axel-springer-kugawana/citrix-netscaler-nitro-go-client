package nitro

type TransformpolicylabelResource struct {
	Labelname       string `json:"labelname"`
	Policylabeltype string `json:"policylabeltype,omitempty"`
}

type TransformpolicylabelKey struct {
	Labelname string
}

type transformpolicylabel_update struct {
	Labelname string `json:"labelname"`
}

type transformpolicylabel_payload struct {
	transformpolicylabel interface{}
}

func transformpolicylabel_key_to_args(key TransformpolicylabelKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteTransformpolicylabel(key TransformpolicylabelKey) error {
	return c.deleteResourceWithArgs("transformpolicylabel", key.Labelname, transformpolicylabel_key_to_args(key))
}

func (c *NitroClient) GetTransformpolicylabel(key TransformpolicylabelKey) (*TransformpolicylabelResource, error) {
	var results struct {
		Transformpolicylabel []TransformpolicylabelResource
	}

	if err := c.getResourceWithArgs("transformpolicylabel", key.Labelname, transformpolicylabel_key_to_args(key), &results); err != nil || len(results.Transformpolicylabel) != 1 {
		return nil, err
	}

	return &results.Transformpolicylabel[0], nil
}

func (c *NitroClient) ListTransformpolicylabel() ([]TransformpolicylabelResource, error) {
	var results struct {
		Transformpolicylabel []TransformpolicylabelResource
	}

	if err := c.listResources("transformpolicylabel", &results); err != nil {
		return nil, err
	}

	return results.Transformpolicylabel, nil
}

func (c *NitroClient) AddTransformpolicylabel(resource TransformpolicylabelResource) error {
	return c.addResource("transformpolicylabel", resource)
}

func (c *NitroClient) RenameTransformpolicylabel(labelname string, newName string) error {
	return c.renameResource("transformpolicylabel", "labelname", labelname, newName)
}
