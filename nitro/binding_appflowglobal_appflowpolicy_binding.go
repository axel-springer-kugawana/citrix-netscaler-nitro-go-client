package nitro

type AppflowglobalAppflowpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
	Type                   string `json:"type,omitempty"`
}

type AppflowglobalAppflowpolicyBindingKey struct {
	Labelname  string
	Policyname string
}

func (c *NitroClient) AddAppflowglobalAppflowpolicyBinding(binding AppflowglobalAppflowpolicyBinding) error {
	return nil
}

func (c *NitroClient) ListAppflowglobalAppflowpolicyBinding() ([]AppflowglobalAppflowpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetAppflowglobalAppflowpolicyBinding(key AppflowglobalAppflowpolicyBindingKey) (*AppflowglobalAppflowpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteAppflowglobalAppflowpolicyBinding(key AppflowglobalAppflowpolicyBindingKey) error {
	return nil
}
