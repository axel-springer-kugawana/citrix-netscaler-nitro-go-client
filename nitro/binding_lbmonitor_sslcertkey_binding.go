package nitro

import (
	"fmt"
	"strconv"
)

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

type get_lbmonitor_sslcertkey_binding struct {
	Results []LbmonitorSslcertkeyBinding `json:"lbmonitor_sslcertkey_binding"`
}

type add_lbmonitor_sslcertkey_binding_payload struct {
	lbmonitor_sslcertkey_binding LbmonitorSslcertkeyBinding
}

func lbmonitor_sslcertkey_binding_key_to_id_args(key LbmonitorSslcertkeyBindingKey) (string, string) {
	var _ = strconv.Itoa

	result := ""

	result = result + ",monitorname:" + key.Monitorname
	result = result + ",certkeyname:" + key.Certkeyname
	return "", result
}

func (c *NitroClient) AddLbmonitorSslcertkeyBinding(binding LbmonitorSslcertkeyBinding) error {
	payload := add_lbmonitor_sslcertkey_binding_payload{
		binding,
	}

	return c.put("lbmonitor_sslcertkey_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbmonitorSslcertkeyBinding(key LbmonitorSslcertkeyBindingKey) ([]LbmonitorSslcertkeyBinding, error) {
	var results get_lbmonitor_sslcertkey_binding

	id, args := lbmonitor_sslcertkey_binding_key_to_id_args(key)

	if err := c.get("lbmonitor_sslcertkey_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbmonitorSslcertkeyBinding() ([]LbmonitorSslcertkeyBinding, error) {
	var results get_lbmonitor_sslcertkey_binding

	if err := c.get("lbmonitor_sslcertkey_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbmonitorSslcertkeyBinding(key LbmonitorSslcertkeyBindingKey) (*LbmonitorSslcertkeyBinding, error) {
	var results get_lbmonitor_sslcertkey_binding

	id, args := lbmonitor_sslcertkey_binding_key_to_id_args(key)

	if err := c.get("lbmonitor_sslcertkey_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbmonitor_sslcertkey_binding element found")
		} else if len(results.Results) < 1 {
			//                        return nil, fmt.Errorf("lbmonitor_sslcertkey_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbmonitorSslcertkeyBinding(key LbmonitorSslcertkeyBindingKey) error {
	id, args := lbmonitor_sslcertkey_binding_key_to_id_args(key)

	return c.delete("lbmonitor_sslcertkey_binding", id, "", args)
}
