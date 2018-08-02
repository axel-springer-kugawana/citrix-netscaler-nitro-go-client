package nitro

type AuthorizationpolicylabelAuthorizationpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Invoke_labelname       string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type AuthorizationpolicylabelAuthorizationpolicyBindingKey struct {
	Labelname  string
	Policyname string
}

func (c *NitroClient) AddAuthorizationpolicylabelAuthorizationpolicyBinding(binding AuthorizationpolicylabelAuthorizationpolicyBinding) error {
	return nil
}

func (c *NitroClient) ListAuthorizationpolicylabelAuthorizationpolicyBinding() ([]AuthorizationpolicylabelAuthorizationpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetAuthorizationpolicylabelAuthorizationpolicyBinding(key AuthorizationpolicylabelAuthorizationpolicyBindingKey) (*AuthorizationpolicylabelAuthorizationpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteAuthorizationpolicylabelAuthorizationpolicyBinding(key AuthorizationpolicylabelAuthorizationpolicyBindingKey) error {
	return nil
}
