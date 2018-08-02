package nitro

type LbvserverVideooptimizationpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverVideooptimizationpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

func (c *NitroClient) AddLbvserverVideooptimizationpolicyBinding(binding LbvserverVideooptimizationpolicyBinding) error {
	return nil
}

func (c *NitroClient) ListLbvserverVideooptimizationpolicyBinding() ([]LbvserverVideooptimizationpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbvserverVideooptimizationpolicyBinding(key LbvserverVideooptimizationpolicyBindingKey) (*LbvserverVideooptimizationpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbvserverVideooptimizationpolicyBinding(key LbvserverVideooptimizationpolicyBindingKey) error {
	return nil
}
