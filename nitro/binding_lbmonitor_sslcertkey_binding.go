package nitro

import (
	"fmt"
	"strconv"
	"strings"
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

type add_lbmonitor_sslcertkey_binding_payload struct {
	Resource LbmonitorSslcertkeyBinding `json:"lbmonitor_sslcertkey_binding"`
}

type get_lbmonitor_sslcertkey_binding_result struct {
	Results []LbmonitorSslcertkeyBinding `json:"lbmonitor_sslcertkey_binding"`
}

type count_lbmonitor_sslcertkey_binding_result struct {
	Results []Count `json:"lbmonitor_sslcertkey_binding"`
}

func lbmonitor_sslcertkey_binding_key_to_id_qs(key LbmonitorSslcertkeyBindingKey, arg string) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "monitorname:"+key.Monitorname)
	args = append(args, "certkeyname:"+key.Certkeyname)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[arg] = strings.Join(args, ",")
	}

	return "", qs
}

func lbmonitor_sslcertkey_binding_key_to_id_args(key LbmonitorSslcertkeyBindingKey) (string, map[string]string) {
	return lbmonitor_sslcertkey_binding_key_to_id_qs(key, "args")
}

func lbmonitor_sslcertkey_binding_key_to_id_filter(key LbmonitorSslcertkeyBindingKey) (string, map[string]string) {
	return lbmonitor_sslcertkey_binding_key_to_id_qs(key, "filter")
}

func (c *NitroClient) AddLbmonitorSslcertkeyBinding(binding LbmonitorSslcertkeyBinding) error {
	payload := add_lbmonitor_sslcertkey_binding_payload{
		binding,
	}

	return c.put("lbmonitor_sslcertkey_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountLbmonitorSslcertkeyBinding() (int, error) {
	var results count_lbmonitor_sslcertkey_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("lbmonitor_sslcertkey_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountLbmonitorSslcertkeyBinding(id string) (int, error) {
	var results count_lbmonitor_sslcertkey_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbmonitor_sslcertkey_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbmonitorSslcertkeyBinding(key LbmonitorSslcertkeyBindingKey) (bool, error) {
	var results count_lbmonitor_sslcertkey_binding_result

	id, qs := lbmonitor_sslcertkey_binding_key_to_id_filter(key)

	qs["count"] = "yes"

	if err := c.get("lbmonitor_sslcertkey_binding", id, qs, &results); err != nil {
		return false, err
	} else {
		if len(results.Results) > 1 {
			return false, fmt.Errorf("More than one lbmonitor_sslcertkey_binding element found")
		}

		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) BulkListLbmonitorSslcertkeyBinding() ([]LbmonitorSslcertkeyBinding, error) {
	var results get_lbmonitor_sslcertkey_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbmonitor_sslcertkey_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListLbmonitorSslcertkeyBinding(id string) ([]LbmonitorSslcertkeyBinding, error) {
	var results get_lbmonitor_sslcertkey_binding_result

	if err := c.get("lbmonitor_sslcertkey_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbmonitorSslcertkeyBinding(key LbmonitorSslcertkeyBindingKey) (*LbmonitorSslcertkeyBinding, error) {
	var results get_lbmonitor_sslcertkey_binding_result

	id, qs := lbmonitor_sslcertkey_binding_key_to_id_args(key)

	if err := c.get("lbmonitor_sslcertkey_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbmonitor_sslcertkey_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbmonitor_sslcertkey_binding element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbmonitorSslcertkeyBinding(key LbmonitorSslcertkeyBindingKey) error {
	id, qs := lbmonitor_sslcertkey_binding_key_to_id_args(key)

	return c.delete("lbmonitor_sslcertkey_binding", id, qs)
}
