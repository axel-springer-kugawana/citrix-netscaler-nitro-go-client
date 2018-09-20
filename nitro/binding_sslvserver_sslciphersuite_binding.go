package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type SslvserverSslciphersuiteBinding struct {
	Ciphername  string `json:"ciphername,omitempty"`
	Vservername string `json:"vservername,omitempty"`
}

type SslvserverSslciphersuiteBindingKey struct {
	Vservername string
	Ciphername  string
}

func (resource SslvserverSslciphersuiteBinding) ToKey() SslvserverSslciphersuiteBindingKey {
	key := SslvserverSslciphersuiteBindingKey{
		resource.Vservername,
		resource.Ciphername,
	}

	return key
}

func (key SslvserverSslciphersuiteBindingKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = url.QueryEscape(key.Vservername)
	args = append(args, "ciphername:"+url.QueryEscape(key.Ciphername))

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key SslvserverSslciphersuiteBindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key SslvserverSslciphersuiteBindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_sslvserver_sslciphersuite_binding_payload struct {
	Resource SslvserverSslciphersuiteBinding `json:"sslvserver_sslciphersuite_binding"`
}

func (c *NitroClient) AddSslvserverSslciphersuiteBinding(resource SslvserverSslciphersuiteBinding) error {
	payload := add_sslvserver_sslciphersuite_binding_payload{
		resource,
	}

	return c.put("sslvserver_sslciphersuite_binding", "", nil, payload)
}

//      LIST

type list_sslvserver_sslciphersuite_binding_result struct {
	Results []SslvserverSslciphersuiteBinding `json:"sslvserver_sslciphersuite_binding"`
}

func (c *NitroClient) ListSslvserverSslciphersuiteBinding() ([]SslvserverSslciphersuiteBinding, error) {
	results := list_sslvserver_sslciphersuite_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("sslvserver_sslciphersuite_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_sslvserver_sslciphersuite_binding_result struct {
	Results []Count `json:"sslvserver_sslciphersuite_binding"`
}

func (c *NitroClient) CountSslvserverSslciphersuiteBinding() (int, error) {
	results := count_sslvserver_sslciphersuite_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("sslvserver_sslciphersuite_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      READ

type get_sslvserver_sslciphersuite_binding_result struct {
	Results []SslvserverSslciphersuiteBinding `json:"sslvserver_sslciphersuite_binding"`
}

func (c *NitroClient) GetSslvserverSslciphersuiteBinding(key SslvserverSslciphersuiteBindingKey) (*SslvserverSslciphersuiteBinding, error) {
	var results get_sslvserver_sslciphersuite_binding_result

	id, qs := key.to_id_filter()

	if err := c.get("sslvserver_sslciphersuite_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one sslvserver_sslciphersuite_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("sslvserver_sslciphersuite_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

func (c *NitroClient) ExistsSslvserverSslciphersuiteBinding(key SslvserverSslciphersuiteBindingKey) (bool, error) {
	var results count_sslvserver_sslciphersuite_binding_result

	id, qs := key.to_id_filter()

	qs["count"] = "yes"

	if err := c.get("sslvserver_sslciphersuite_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteSslvserverSslciphersuiteBinding(key SslvserverSslciphersuiteBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("sslvserver_sslciphersuite_binding", id, qs)
}
