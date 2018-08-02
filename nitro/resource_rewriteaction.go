package nitro

type Rewriteaction struct {
	Name              string `json:"name"`
	Bypasssafetycheck string `json:"bypasssafetycheck,omitempty"`
	Comment           string `json:"comment,omitempty"`
	Pattern           string `json:"pattern,omitempty"`
	Refinesearch      string `json:"refinesearch,omitempty"`
	Search            string `json:"search,omitempty"`
	Stringbuilderexpr string `json:"stringbuilderexpr,omitempty"`
	Target            string `json:"target,omitempty"`
	Type              string `json:"type,omitempty"`
}

type RewriteactionKey struct {
	Name string
}

type rewriteaction_update struct {
	Name              string `json:"name"`
	Target            string `json:"target,omitempty"`
	Stringbuilderexpr string `json:"stringbuilderexpr,omitempty"`
	Pattern           string `json:"pattern,omitempty"`
	Search            string `json:"search,omitempty"`
	Bypasssafetycheck string `json:"bypasssafetycheck,omitempty"`
	Refinesearch      string `json:"refinesearch,omitempty"`
	Comment           string `json:"comment,omitempty"`
}

type rewriteaction_payload struct {
	rewriteaction interface{}
}

func rewriteaction_key_to_args(key RewriteactionKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteRewriteaction(key RewriteactionKey) error {
	return c.deleteResourceWithArgs("rewriteaction", key.Name, rewriteaction_key_to_args(key))
}

func (c *NitroClient) GetRewriteaction(key RewriteactionKey) (*Rewriteaction, error) {
	var results struct {
		Rewriteaction []Rewriteaction
	}

	if err := c.getResourceWithArgs("rewriteaction", key.Name, rewriteaction_key_to_args(key), &results); err != nil || len(results.Rewriteaction) != 1 {
		return nil, err
	}

	return &results.Rewriteaction[0], nil
}

func (c *NitroClient) ListRewriteaction() ([]Rewriteaction, error) {
	var results struct {
		Rewriteaction []Rewriteaction
	}

	if err := c.listResources("rewriteaction", &results); err != nil {
		return nil, err
	}

	return results.Rewriteaction, nil
}

func (c *NitroClient) AddRewriteaction(resource Rewriteaction) error {
	return c.addResource("rewriteaction", resource)
}

func (c *NitroClient) RenameRewriteaction(name string, newName string) error {
	return c.renameResource("rewriteaction", "name", name, newName)
}

func (c *NitroClient) UnsetRewriteaction(name string, fields ...string) error {
	return c.unsetResource("rewriteaction", "name", name, fields)
}

func (c *NitroClient) UpdateRewriteaction(resource Rewriteaction) error {
	update := rewriteaction_update{
		resource.Name,
		resource.Target,
		resource.Stringbuilderexpr,
		resource.Pattern,
		resource.Search,
		resource.Bypasssafetycheck,
		resource.Refinesearch,
		resource.Comment,
	}

	return c.Put("rewriteaction", update)
}
