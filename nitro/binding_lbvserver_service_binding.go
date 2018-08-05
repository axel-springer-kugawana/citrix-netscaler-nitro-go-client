package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbvserverServiceBinding struct {
	Name        string `json:"name,omitempty"`
	Servicename string `json:"servicename,omitempty"`
	Weight      int    `json:"weight,string,omitempty"`
}

type LbvserverServiceBindingKey struct {
	Name        string
	Servicename string
}

type add_lbvserver_service_binding_payload struct {
	Resource LbvserverServiceBinding `json:"lbvserver_service_binding"`
}

type get_lbvserver_service_binding_result struct {
	Results []LbvserverServiceBinding `json:"lbvserver_service_binding"`
}

type count_lbvserver_service_binding_result struct {
	Results []Count `json:"lbvserver_service_binding"`
}

func lbvserver_service_binding_key_to_id_qs(key LbvserverServiceBindingKey, arg string) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)
	args = append(args, "servicename:"+key.Servicename)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[arg] = strings.Join(args, ",")
	}

	return "", qs
}

func lbvserver_service_binding_key_to_id_args(key LbvserverServiceBindingKey) (string, map[string]string) {
	return lbvserver_service_binding_key_to_id_qs(key, "args")
}

func lbvserver_service_binding_key_to_id_filter(key LbvserverServiceBindingKey) (string, map[string]string) {
	return lbvserver_service_binding_key_to_id_qs(key, "filter")
}

func (c *NitroClient) AddLbvserverServiceBinding(binding LbvserverServiceBinding) error {
	payload := add_lbvserver_service_binding_payload{
		binding,
	}

	return c.put("lbvserver_service_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountLbvserverServiceBinding() (int, error) {
	var results count_lbvserver_service_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("lbvserver_service_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountLbvserverServiceBinding(id string) (int, error) {
	var results count_lbvserver_service_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbvserver_service_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbvserverServiceBinding(key LbvserverServiceBindingKey) (bool, error) {
	var results count_lbvserver_service_binding_result

	id, qs := lbvserver_service_binding_key_to_id_filter(key)

	qs["count"] = "yes"

	if err := c.get("lbvserver_service_binding", id, qs, &results); err != nil {
		return false, err
	} else {
		if len(results.Results) > 1 {
			return false, fmt.Errorf("More than one lbvserver_service_binding element found")
		}

		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) BulkListLbvserverServiceBinding() ([]LbvserverServiceBinding, error) {
	var results get_lbvserver_service_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_service_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListLbvserverServiceBinding(id string) ([]LbvserverServiceBinding, error) {
	var results get_lbvserver_service_binding_result

	if err := c.get("lbvserver_service_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverServiceBinding(key LbvserverServiceBindingKey) (*LbvserverServiceBinding, error) {
	var results get_lbvserver_service_binding_result

	id, qs := lbvserver_service_binding_key_to_id_args(key)

	if err := c.get("lbvserver_service_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_service_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbvserver_service_binding element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverServiceBinding(key LbvserverServiceBindingKey) error {
	id, qs := lbvserver_service_binding_key_to_id_args(key)

	return c.delete("lbvserver_service_binding", id, qs)
}
