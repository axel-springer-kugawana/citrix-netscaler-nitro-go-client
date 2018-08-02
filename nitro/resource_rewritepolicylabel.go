package nitro

type RewritepolicylabelResource struct {
	Labelname string `json:"labelname"`
	Comment   string `json:"comment,omitempty"`
	Transform string `json:"transform,omitempty"`
}

type RewritepolicylabelKey struct {
	Labelname string
}

type rewritepolicylabel_update struct {
	Labelname string `json:"labelname"`
}

type rewritepolicylabel_payload struct {
	rewritepolicylabel interface{}
}

func rewritepolicylabel_key_to_args(key RewritepolicylabelKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteRewritepolicylabel(key RewritepolicylabelKey) error {
	return c.deleteResourceWithArgs("rewritepolicylabel", key.Labelname, rewritepolicylabel_key_to_args(key))
}

func (c *NitroClient) GetRewritepolicylabel(key RewritepolicylabelKey) (*RewritepolicylabelResource, error) {
	var results struct {
		Rewritepolicylabel []RewritepolicylabelResource
	}

	if err := c.getResourceWithArgs("rewritepolicylabel", key.Labelname, rewritepolicylabel_key_to_args(key), &results); err != nil || len(results.Rewritepolicylabel) != 1 {
		return nil, err
	}

	return &results.Rewritepolicylabel[0], nil
}

func (c *NitroClient) ListRewritepolicylabel() ([]RewritepolicylabelResource, error) {
	var results struct {
		Rewritepolicylabel []RewritepolicylabelResource
	}

	if err := c.listResources("rewritepolicylabel", &results); err != nil {
		return nil, err
	}

	return results.Rewritepolicylabel, nil
}

func (c *NitroClient) AddRewritepolicylabel(resource RewritepolicylabelResource) error {
	return c.addResource("rewritepolicylabel", resource)
}

func (c *NitroClient) RenameRewritepolicylabel(labelname string, newName string) error {
	return c.renameResource("rewritepolicylabel", "labelname", labelname, newName)
}
