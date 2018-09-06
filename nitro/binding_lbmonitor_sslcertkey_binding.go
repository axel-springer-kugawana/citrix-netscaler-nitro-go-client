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

func (resource LbmonitorSslcertkeyBinding) ToKey() LbmonitorSslcertkeyBindingKey {
	key := LbmonitorSslcertkeyBindingKey{
		resource.Monitorname,
		resource.Certkeyname,
	}

	return key
}

func (key LbmonitorSslcertkeyBindingKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = key.Monitorname
	args = append(args, "certkeyname:"+key.Certkeyname)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key LbmonitorSslcertkeyBindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key LbmonitorSslcertkeyBindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_lbmonitor_sslcertkey_binding_payload struct {
	Resource LbmonitorSslcertkeyBinding `json:"lbmonitor_sslcertkey_binding"`
}

func (c *NitroClient) AddLbmonitorSslcertkeyBinding(resource LbmonitorSslcertkeyBinding) error {
	payload := add_lbmonitor_sslcertkey_binding_payload{
		resource,
	}

	return c.put("lbmonitor_sslcertkey_binding", "", nil, payload)
}

//      LIST

type list_lbmonitor_sslcertkey_binding_result struct {
	Results []LbmonitorSslcertkeyBinding `json:"lbmonitor_sslcertkey_binding"`
}

func (c *NitroClient) ListLbmonitorSslcertkeyBinding() ([]LbmonitorSslcertkeyBinding, error) {
	results := list_lbmonitor_sslcertkey_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbmonitor_sslcertkey_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_lbmonitor_sslcertkey_binding_result struct {
	Results []Count `json:"lbmonitor_sslcertkey_binding"`
}

func (c *NitroClient) CountLbmonitorSslcertkeyBinding() (int, error) {
	results := count_lbmonitor_sslcertkey_binding_result{}

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

//      READ

type get_lbmonitor_sslcertkey_binding_result struct {
	Results []LbmonitorSslcertkeyBinding `json:"lbmonitor_sslcertkey_binding"`
}

func (c *NitroClient) GetLbmonitorSslcertkeyBinding(key LbmonitorSslcertkeyBindingKey) (*LbmonitorSslcertkeyBinding, error) {
	var results get_lbmonitor_sslcertkey_binding_result

	id, qs := key.to_id_filter()

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

//      EXISTS

func (c *NitroClient) ExistsLbmonitorSslcertkeyBinding(key LbmonitorSslcertkeyBindingKey) (bool, error) {
	var results count_lbmonitor_sslcertkey_binding_result

	id, qs := key.to_id_filter()

	qs["count"] = "yes"

	if err := c.get("lbmonitor_sslcertkey_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteLbmonitorSslcertkeyBinding(key LbmonitorSslcertkeyBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("lbmonitor_sslcertkey_binding", id, qs)
}
