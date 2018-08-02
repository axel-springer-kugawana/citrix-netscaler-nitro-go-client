package nitro

type LbmonitorSslcertkeyBinding struct {
	Ca          bool   `json:"ca,omitempty"`
	Certkeyname string `json:"certkeyname,omitempty"`
	Crlcheck    string `json:"crlcheck,omitempty"`
	Monitorname string `json:"monitorname,omitempty"`
	Ocspcheck   string `json:"ocspcheck,omitempty"`
}

type LbmonitorSslcertkeyBindingKey struct {
	Monitorname string
	Certkeyname string
}

func (c *NitroClient) AddLbmonitorSslcertkeyBinding(binding LbmonitorSslcertkeyBinding) error {
	return nil
}

func (c *NitroClient) ListLbmonitorSslcertkeyBinding() ([]LbmonitorSslcertkeyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbmonitorSslcertkeyBinding(key LbmonitorSslcertkeyBindingKey) (*LbmonitorSslcertkeyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbmonitorSslcertkeyBinding(key LbmonitorSslcertkeyBindingKey) error {
	return nil
}
