package nitro

import (
	"fmt"
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

type add_lbvserver_servicegroup_binding_payload struct {
	Resource LbvserverServicegroupBinding `json:"lbvserver_servicegroup_binding"`
}

type get_lbvserver_servicegroup_binding_result struct {
	Results []LbvserverServicegroupBinding `json:"lbvserver_servicegroup_binding"`
}

type count_lbvserver_servicegroup_binding_result struct {
	Results []Count `json:"lbvserver_servicegroup_binding"`
}

func lbvserver_servicegroup_binding_key_to_id_qs(key LbvserverServicegroupBindingKey, arg string) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)
	args = append(args, "servicegroupname:"+key.Servicegroupname)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[arg] = strings.Join(args, ",")
	}

	return "", qs
}

func lbvserver_servicegroup_binding_key_to_id_args(key LbvserverServicegroupBindingKey) (string, map[string]string) {
	return lbvserver_servicegroup_binding_key_to_id_qs(key, "args")
}

func lbvserver_servicegroup_binding_key_to_id_filter(key LbvserverServicegroupBindingKey) (string, map[string]string) {
	return lbvserver_servicegroup_binding_key_to_id_qs(key, "filter")
}

func (c *NitroClient) AddLbvserverServicegroupBinding(binding LbvserverServicegroupBinding) error {
	payload := add_lbvserver_servicegroup_binding_payload{
		binding,
	}

	return c.put("lbvserver_servicegroup_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountLbvserverServicegroupBinding() (int, error) {
	var results count_lbvserver_servicegroup_binding_result

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

func (c *NitroClient) CountLbvserverServicegroupBinding(id string) (int, error) {
	var results count_lbvserver_servicegroup_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbvserver_servicegroup_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbvserverServicegroupBinding(key LbvserverServicegroupBindingKey) (bool, error) {
	var results count_lbvserver_servicegroup_binding_result

	id, qs := lbvserver_servicegroup_binding_key_to_id_filter(key)

	qs["count"] = "yes"

	if err := c.get("lbvserver_servicegroup_binding", id, qs, &results); err != nil {
		return false, err
	} else {
		if len(results.Results) > 1 {
			return false, fmt.Errorf("More than one lbvserver_servicegroup_binding element found")
		}

		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) BulkListLbvserverServicegroupBinding() ([]LbvserverServicegroupBinding, error) {
	var results get_lbvserver_servicegroup_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_servicegroup_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListLbvserverServicegroupBinding(id string) ([]LbvserverServicegroupBinding, error) {
	var results get_lbvserver_servicegroup_binding_result

	if err := c.get("lbvserver_servicegroup_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverServicegroupBinding(key LbvserverServicegroupBindingKey) (*LbvserverServicegroupBinding, error) {
	var results get_lbvserver_servicegroup_binding_result

	id, qs := lbvserver_servicegroup_binding_key_to_id_args(key)

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

func (c *NitroClient) DeleteLbvserverServicegroupBinding(key LbvserverServicegroupBindingKey) error {
	id, qs := lbvserver_servicegroup_binding_key_to_id_args(key)

	return c.delete("lbvserver_servicegroup_binding", id, qs)
}
