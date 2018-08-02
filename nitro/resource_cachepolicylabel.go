package nitro

type Cachepolicylabel struct {
	Labelname string `json:"labelname"`
	Evaluates string `json:"evaluates,omitempty"`
}

type CachepolicylabelKey struct {
	Labelname string
}

type cachepolicylabel_update struct {
	Labelname string `json:"labelname"`
}

type cachepolicylabel_payload struct {
	cachepolicylabel interface{}
}

func cachepolicylabel_key_to_args(key CachepolicylabelKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteCachepolicylabel(key CachepolicylabelKey) error {
	return c.deleteResourceWithArgs("cachepolicylabel", key.Labelname, cachepolicylabel_key_to_args(key))
}

func (c *NitroClient) GetCachepolicylabel(key CachepolicylabelKey) (*Cachepolicylabel, error) {
	var results struct {
		Cachepolicylabel []Cachepolicylabel
	}

	if err := c.getResourceWithArgs("cachepolicylabel", key.Labelname, cachepolicylabel_key_to_args(key), &results); err != nil || len(results.Cachepolicylabel) != 1 {
		return nil, err
	}

	return &results.Cachepolicylabel[0], nil
}

func (c *NitroClient) ListCachepolicylabel() ([]Cachepolicylabel, error) {
	var results struct {
		Cachepolicylabel []Cachepolicylabel
	}

	if err := c.listResources("cachepolicylabel", &results); err != nil {
		return nil, err
	}

	return results.Cachepolicylabel, nil
}

func (c *NitroClient) AddCachepolicylabel(resource Cachepolicylabel) error {
	return c.addResource("cachepolicylabel", resource)
}

func (c *NitroClient) RenameCachepolicylabel(labelname string, newName string) error {
	return c.renameResource("cachepolicylabel", "labelname", labelname, newName)
}
