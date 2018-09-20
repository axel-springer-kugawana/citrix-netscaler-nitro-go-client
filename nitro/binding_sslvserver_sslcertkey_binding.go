package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type SslvserverSslcertkeyBinding struct {
	Ca          bool   `json:"ca,omitempty"`
	Certkeyname string `json:"certkeyname,omitempty"`
	Crlcheck    string `json:"crlcheck,omitempty"`
	Ocspcheck   string `json:"ocspcheck,omitempty"`
	Skipcaname  bool   `json:"skipcaname,omitempty"`
	Snicert     bool   `json:"snicert,omitempty"`
	Vservername string `json:"vservername,omitempty"`
}

type SslvserverSslcertkeyBindingKey struct {
	Vservername string
	Certkeyname string
	Ca          bool
}

func (resource SslvserverSslcertkeyBinding) ToKey() SslvserverSslcertkeyBindingKey {
	key := SslvserverSslcertkeyBindingKey{
		resource.Vservername,
		resource.Certkeyname,
		resource.Ca,
	}

	return key
}

func (key SslvserverSslcertkeyBindingKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = url.QueryEscape(key.Vservername)
	args = append(args, "certkeyname:"+url.QueryEscape(key.Certkeyname))
	args = append(args, "ca:"+strconv.FormatBool(key.Ca))

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key SslvserverSslcertkeyBindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key SslvserverSslcertkeyBindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_sslvserver_sslcertkey_binding_payload struct {
	Resource SslvserverSslcertkeyBinding `json:"sslvserver_sslcertkey_binding"`
}

func (c *NitroClient) AddSslvserverSslcertkeyBinding(resource SslvserverSslcertkeyBinding) error {
	payload := add_sslvserver_sslcertkey_binding_payload{
		resource,
	}

	return c.put("sslvserver_sslcertkey_binding", "", nil, payload)
}

//      LIST

type list_sslvserver_sslcertkey_binding_result struct {
	Results []SslvserverSslcertkeyBinding `json:"sslvserver_sslcertkey_binding"`
}

func (c *NitroClient) ListSslvserverSslcertkeyBinding() ([]SslvserverSslcertkeyBinding, error) {
	results := list_sslvserver_sslcertkey_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("sslvserver_sslcertkey_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_sslvserver_sslcertkey_binding_result struct {
	Results []Count `json:"sslvserver_sslcertkey_binding"`
}

func (c *NitroClient) CountSslvserverSslcertkeyBinding() (int, error) {
	results := count_sslvserver_sslcertkey_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("sslvserver_sslcertkey_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      READ

type get_sslvserver_sslcertkey_binding_result struct {
	Results []SslvserverSslcertkeyBinding `json:"sslvserver_sslcertkey_binding"`
}

func (c *NitroClient) GetSslvserverSslcertkeyBinding(key SslvserverSslcertkeyBindingKey) (*SslvserverSslcertkeyBinding, error) {
	var results get_sslvserver_sslcertkey_binding_result

	id, qs := key.to_id_filter()

	if err := c.get("sslvserver_sslcertkey_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one sslvserver_sslcertkey_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("sslvserver_sslcertkey_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

func (c *NitroClient) ExistsSslvserverSslcertkeyBinding(key SslvserverSslcertkeyBindingKey) (bool, error) {
	var results count_sslvserver_sslcertkey_binding_result

	id, qs := key.to_id_filter()

	qs["count"] = "yes"

	if err := c.get("sslvserver_sslcertkey_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteSslvserverSslcertkeyBinding(key SslvserverSslcertkeyBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("sslvserver_sslcertkey_binding", id, qs)
}
