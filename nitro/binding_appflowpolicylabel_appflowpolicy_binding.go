package nitro

type AppflowpolicylabelAppflowpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Invoke_labelname       string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type AppflowpolicylabelAppflowpolicyBindingKey struct {
	Labelname  string
	Policyname string
}

func (c *NitroClient) AddAppflowpolicylabelAppflowpolicyBinding(binding AppflowpolicylabelAppflowpolicyBinding) error {
	return nil
}

func (c *NitroClient) ListAppflowpolicylabelAppflowpolicyBinding() ([]AppflowpolicylabelAppflowpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetAppflowpolicylabelAppflowpolicyBinding(key AppflowpolicylabelAppflowpolicyBindingKey) (*AppflowpolicylabelAppflowpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteAppflowpolicylabelAppflowpolicyBinding(key AppflowpolicylabelAppflowpolicyBindingKey) error {
	return nil
}
