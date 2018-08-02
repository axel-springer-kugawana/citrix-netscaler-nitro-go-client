package nitro

type LbvserverRewritepolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverRewritepolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

func (c *NitroClient) AddLbvserverRewritepolicyBinding(binding LbvserverRewritepolicyBinding) error {
	return nil
}

func (c *NitroClient) ListLbvserverRewritepolicyBinding() ([]LbvserverRewritepolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbvserverRewritepolicyBinding(key LbvserverRewritepolicyBindingKey) (*LbvserverRewritepolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbvserverRewritepolicyBinding(key LbvserverRewritepolicyBindingKey) error {
	return nil
}
