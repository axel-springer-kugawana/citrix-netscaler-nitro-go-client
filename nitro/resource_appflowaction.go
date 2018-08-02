package nitro

type AppflowactionResource struct {
	Name                   string   `json:"name"`
	Clientsidemeasurements string   `json:"clientsidemeasurements,omitempty"`
	Collectors             []string `json:"collectors,omitempty"`
	Comment                string   `json:"comment,omitempty"`
	Distributionalgorithm  string   `json:"distributionalgorithm,omitempty"`
	Metricslog             bool     `json:"metricslog,omitempty"`
	Pagetracking           string   `json:"pagetracking,omitempty"`
	Securityinsight        string   `json:"securityinsight,omitempty"`
	Transactionlog         string   `json:"transactionlog,omitempty"`
	Videoanalytics         string   `json:"videoanalytics,omitempty"`
	Webinsight             string   `json:"webinsight,omitempty"`
}

type AppflowactionKey struct {
	Name string
}

type appflowaction_update struct {
	Name                   string   `json:"name"`
	Collectors             []string `json:"collectors,omitempty"`
	Clientsidemeasurements string   `json:"clientsidemeasurements,omitempty"`
	Comment                string   `json:"comment,omitempty"`
	Pagetracking           string   `json:"pagetracking,omitempty"`
	Webinsight             string   `json:"webinsight,omitempty"`
	Securityinsight        string   `json:"securityinsight,omitempty"`
	Videoanalytics         string   `json:"videoanalytics,omitempty"`
	Distributionalgorithm  string   `json:"distributionalgorithm,omitempty"`
}

type appflowaction_payload struct {
	appflowaction interface{}
}

func appflowaction_key_to_args(key AppflowactionKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteAppflowaction(key AppflowactionKey) error {
	return c.deleteResourceWithArgs("appflowaction", key.Name, appflowaction_key_to_args(key))
}

func (c *NitroClient) GetAppflowaction(key AppflowactionKey) (*AppflowactionResource, error) {
	var results struct {
		Appflowaction []AppflowactionResource
	}

	if err := c.getResourceWithArgs("appflowaction", key.Name, appflowaction_key_to_args(key), &results); err != nil || len(results.Appflowaction) != 1 {
		return nil, err
	}

	return &results.Appflowaction[0], nil
}

func (c *NitroClient) ListAppflowaction() ([]AppflowactionResource, error) {
	var results struct {
		Appflowaction []AppflowactionResource
	}

	if err := c.listResources("appflowaction", &results); err != nil {
		return nil, err
	}

	return results.Appflowaction, nil
}

func (c *NitroClient) AddAppflowaction(resource AppflowactionResource) error {
	return c.addResource("appflowaction", resource)
}

func (c *NitroClient) RenameAppflowaction(name string, newName string) error {
	return c.renameResource("appflowaction", "name", name, newName)
}

func (c *NitroClient) UnsetAppflowaction(name string, fields ...string) error {
	return c.unsetResource("appflowaction", "name", name, fields)
}

func (c *NitroClient) UpdateAppflowaction(resource AppflowactionResource) error {
	update := appflowaction_update{
		resource.Name,
		resource.Collectors,
		resource.Clientsidemeasurements,
		resource.Comment,
		resource.Pagetracking,
		resource.Webinsight,
		resource.Securityinsight,
		resource.Videoanalytics,
		resource.Distributionalgorithm,
	}

	return c.Put("appflowaction", update)
}
