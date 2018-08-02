package nitro

type LbvserverTmtrafficpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverTmtrafficpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

func (c *NitroClient) AddLbvserverTmtrafficpolicyBinding(binding LbvserverTmtrafficpolicyBinding) error {
	return nil
}

func (c *NitroClient) ListLbvserverTmtrafficpolicyBinding() ([]LbvserverTmtrafficpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbvserverTmtrafficpolicyBinding(key LbvserverTmtrafficpolicyBindingKey) (*LbvserverTmtrafficpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbvserverTmtrafficpolicyBinding(key LbvserverTmtrafficpolicyBindingKey) error {
	return nil
}
