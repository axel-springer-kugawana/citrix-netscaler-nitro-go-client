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

func (resource LbvserverServiceBinding) ToKey() LbvserverServiceBindingKey {
	key := LbvserverServiceBindingKey{
		resource.Name,
		resource.Servicename,
	}

	return key
}

func (key LbvserverServiceBindingKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = key.Name
	args = append(args, "servicename:"+key.Servicename)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key LbvserverServiceBindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key LbvserverServiceBindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_lbvserver_service_binding_payload struct {
	Resource LbvserverServiceBinding `json:"lbvserver_service_binding"`
}

func (c *NitroClient) AddLbvserverServiceBinding(resource LbvserverServiceBinding) error {
	payload := add_lbvserver_service_binding_payload{
		resource,
	}

	return c.put("lbvserver_service_binding", "", nil, payload)
}

//      LIST

type list_lbvserver_service_binding_result struct {
	Results []LbvserverServiceBinding `json:"lbvserver_service_binding"`
}

func (c *NitroClient) ListLbvserverServiceBinding() ([]LbvserverServiceBinding, error) {
	results := list_lbvserver_service_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_service_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_lbvserver_service_binding_result struct {
	Results []Count `json:"lbvserver_service_binding"`
}

func (c *NitroClient) CountLbvserverServiceBinding() (int, error) {
	results := count_lbvserver_service_binding_result{}

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

//      READ

type get_lbvserver_service_binding_result struct {
	Results []LbvserverServiceBinding `json:"lbvserver_service_binding"`
}

func (c *NitroClient) GetLbvserverServiceBinding(key LbvserverServiceBindingKey) (*LbvserverServiceBinding, error) {
	var results get_lbvserver_service_binding_result

	id, qs := key.to_id_filter()

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

//      EXISTS

func (c *NitroClient) ExistsLbvserverServiceBinding(key LbvserverServiceBindingKey) (bool, error) {
	var results count_lbvserver_service_binding_result

	id, qs := key.to_id_filter()

	qs["count"] = "yes"

	if err := c.get("lbvserver_service_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteLbvserverServiceBinding(key LbvserverServiceBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("lbvserver_service_binding", id, qs)
}
