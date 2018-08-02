package nitro

type LbvserverScpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverScpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

func (c *NitroClient) AddLbvserverScpolicyBinding(binding LbvserverScpolicyBinding) error {
	return nil
}

func (c *NitroClient) ListLbvserverScpolicyBinding() ([]LbvserverScpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbvserverScpolicyBinding(key LbvserverScpolicyBindingKey) (*LbvserverScpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbvserverScpolicyBinding(key LbvserverScpolicyBindingKey) error {
	return nil
}
