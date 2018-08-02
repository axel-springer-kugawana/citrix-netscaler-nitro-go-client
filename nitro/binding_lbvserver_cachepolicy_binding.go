package nitro

type LbvserverCachepolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverCachepolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

func (c *NitroClient) AddLbvserverCachepolicyBinding(binding LbvserverCachepolicyBinding) error {
	return nil
}

func (c *NitroClient) ListLbvserverCachepolicyBinding() ([]LbvserverCachepolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbvserverCachepolicyBinding(key LbvserverCachepolicyBindingKey) (*LbvserverCachepolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbvserverCachepolicyBinding(key LbvserverCachepolicyBindingKey) error {
	return nil
}
