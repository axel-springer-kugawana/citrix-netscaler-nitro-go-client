package nitro

type LbvserverAuthorizationpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverAuthorizationpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

func (c *NitroClient) AddLbvserverAuthorizationpolicyBinding(binding LbvserverAuthorizationpolicyBinding) error {
	return nil
}

func (c *NitroClient) ListLbvserverAuthorizationpolicyBinding() ([]LbvserverAuthorizationpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbvserverAuthorizationpolicyBinding(key LbvserverAuthorizationpolicyBindingKey) (*LbvserverAuthorizationpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbvserverAuthorizationpolicyBinding(key LbvserverAuthorizationpolicyBindingKey) error {
	return nil
}
