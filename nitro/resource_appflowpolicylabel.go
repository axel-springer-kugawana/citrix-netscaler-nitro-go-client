package nitro

type AppflowpolicylabelResource struct {
	Labelname       string `json:"labelname"`
	Policylabeltype string `json:"policylabeltype,omitempty"`
}

type AppflowpolicylabelKey struct {
	Labelname string
}

type appflowpolicylabel_update struct {
	Labelname string `json:"labelname"`
}

type appflowpolicylabel_payload struct {
	appflowpolicylabel interface{}
}

func appflowpolicylabel_key_to_args(key AppflowpolicylabelKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteAppflowpolicylabel(key AppflowpolicylabelKey) error {
	return c.deleteResourceWithArgs("appflowpolicylabel", key.Labelname, appflowpolicylabel_key_to_args(key))
}

func (c *NitroClient) GetAppflowpolicylabel(key AppflowpolicylabelKey) (*AppflowpolicylabelResource, error) {
	var results struct {
		Appflowpolicylabel []AppflowpolicylabelResource
	}

	if err := c.getResourceWithArgs("appflowpolicylabel", key.Labelname, appflowpolicylabel_key_to_args(key), &results); err != nil || len(results.Appflowpolicylabel) != 1 {
		return nil, err
	}

	return &results.Appflowpolicylabel[0], nil
}

func (c *NitroClient) ListAppflowpolicylabel() ([]AppflowpolicylabelResource, error) {
	var results struct {
		Appflowpolicylabel []AppflowpolicylabelResource
	}

	if err := c.listResources("appflowpolicylabel", &results); err != nil {
		return nil, err
	}

	return results.Appflowpolicylabel, nil
}

func (c *NitroClient) AddAppflowpolicylabel(resource AppflowpolicylabelResource) error {
	return c.addResource("appflowpolicylabel", resource)
}

func (c *NitroClient) RenameAppflowpolicylabel(labelname string, newName string) error {
	return c.renameResource("appflowpolicylabel", "labelname", labelname, newName)
}
