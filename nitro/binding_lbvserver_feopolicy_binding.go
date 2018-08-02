package nitro

type LbvserverFeopolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverFeopolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

func (c *NitroClient) AddLbvserverFeopolicyBinding(binding LbvserverFeopolicyBinding) error {
	return nil
}

func (c *NitroClient) ListLbvserverFeopolicyBinding() ([]LbvserverFeopolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbvserverFeopolicyBinding(key LbvserverFeopolicyBindingKey) (*LbvserverFeopolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbvserverFeopolicyBinding(key LbvserverFeopolicyBindingKey) error {
	return nil
}
