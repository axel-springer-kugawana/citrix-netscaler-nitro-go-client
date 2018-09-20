package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type LbvserverServicegroupBinding struct {
	Name             string `json:"name,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Weight           int    `json:"weight,string,omitempty"`
}

type LbvserverServicegroupBindingKey struct {
	Name             string
	Servicegroupname string
}

func (resource LbvserverServicegroupBinding) ToKey() LbvserverServicegroupBindingKey {
	key := LbvserverServicegroupBindingKey{
		resource.Name,
		resource.Servicegroupname,
	}

	return key
}

func (key LbvserverServicegroupBindingKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = url.QueryEscape(key.Name)
	args = append(args, "servicegroupname:"+url.QueryEscape(key.Servicegroupname))

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key LbvserverServicegroupBindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key LbvserverServicegroupBindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_lbvserver_servicegroup_binding_payload struct {
	Resource LbvserverServicegroupBinding `json:"lbvserver_servicegroup_binding"`
}

func (c *NitroClient) AddLbvserverServicegroupBinding(resource LbvserverServicegroupBinding) error {
	payload := add_lbvserver_servicegroup_binding_payload{
		resource,
	}

	return c.put("lbvserver_servicegroup_binding", "", nil, payload)
}

//      LIST

type list_lbvserver_servicegroup_binding_result struct {
	Results []LbvserverServicegroupBinding `json:"lbvserver_servicegroup_binding"`
}

func (c *NitroClient) ListLbvserverServicegroupBinding() ([]LbvserverServicegroupBinding, error) {
	results := list_lbvserver_servicegroup_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_servicegroup_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_lbvserver_servicegroup_binding_result struct {
	Results []Count `json:"lbvserver_servicegroup_binding"`
}

func (c *NitroClient) CountLbvserverServicegroupBinding() (int, error) {
	results := count_lbvserver_servicegroup_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("lbvserver_servicegroup_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      READ

type get_lbvserver_servicegroup_binding_result struct {
	Results []LbvserverServicegroupBinding `json:"lbvserver_servicegroup_binding"`
}

func (c *NitroClient) GetLbvserverServicegroupBinding(key LbvserverServicegroupBindingKey) (*LbvserverServicegroupBinding, error) {
	var results get_lbvserver_servicegroup_binding_result

	id, qs := key.to_id_filter()

	if err := c.get("lbvserver_servicegroup_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_servicegroup_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbvserver_servicegroup_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

func (c *NitroClient) ExistsLbvserverServicegroupBinding(key LbvserverServicegroupBindingKey) (bool, error) {
	var results count_lbvserver_servicegroup_binding_result

	id, qs := key.to_id_filter()

	qs["count"] = "yes"

	if err := c.get("lbvserver_servicegroup_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteLbvserverServicegroupBinding(key LbvserverServicegroupBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("lbvserver_servicegroup_binding", id, qs)
}
