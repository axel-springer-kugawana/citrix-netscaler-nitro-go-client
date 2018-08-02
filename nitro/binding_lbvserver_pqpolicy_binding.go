package nitro

type LbvserverPqpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverPqpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

func (c *NitroClient) AddLbvserverPqpolicyBinding(binding LbvserverPqpolicyBinding) error {
	return nil
}

func (c *NitroClient) ListLbvserverPqpolicyBinding() ([]LbvserverPqpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbvserverPqpolicyBinding(key LbvserverPqpolicyBindingKey) (*LbvserverPqpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbvserverPqpolicyBinding(key LbvserverPqpolicyBindingKey) error {
	return nil
}
