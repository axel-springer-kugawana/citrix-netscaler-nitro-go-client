package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type ServicegroupServicegroupmemberBinding struct {
	Port             int    `json:"port,omitempty"`
	Servername       string `json:"servername,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Weight           int    `json:"weight,string,omitempty"`
}

type ServicegroupServicegroupmemberBindingKey struct {
	Servicegroupname string
	Servername       string
	Port             int
}

type add_servicegroup_servicegroupmember_binding_payload struct {
	Resource ServicegroupServicegroupmemberBinding `json:"servicegroup_servicegroupmember_binding"`
}

type get_servicegroup_servicegroupmember_binding_result struct {
	Results []ServicegroupServicegroupmemberBinding `json:"servicegroup_servicegroupmember_binding"`
}

type count_servicegroup_servicegroupmember_binding_result struct {
	Results []Count `json:"servicegroup_servicegroupmember_binding"`
}

func servicegroup_servicegroupmember_binding_key_to_id_qs(key ServicegroupServicegroupmemberBindingKey, arg string) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "servicegroupname:"+key.Servicegroupname)
	args = append(args, "servername:"+key.Servername)
	args = append(args, "port:"+strconv.Itoa(key.Port))

	qs := map[string]string{}

	if len(args) > 0 {
		qs[arg] = strings.Join(args, ",")
	}

	return "", qs
}

func servicegroup_servicegroupmember_binding_key_to_id_args(key ServicegroupServicegroupmemberBindingKey) (string, map[string]string) {
	return servicegroup_servicegroupmember_binding_key_to_id_qs(key, "args")
}

func servicegroup_servicegroupmember_binding_key_to_id_filter(key ServicegroupServicegroupmemberBindingKey) (string, map[string]string) {
	return servicegroup_servicegroupmember_binding_key_to_id_qs(key, "filter")
}

func (c *NitroClient) AddServicegroupServicegroupmemberBinding(binding ServicegroupServicegroupmemberBinding) error {
	payload := add_servicegroup_servicegroupmember_binding_payload{
		binding,
	}

	return c.put("servicegroup_servicegroupmember_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountServicegroupServicegroupmemberBinding() (int, error) {
	var results count_servicegroup_servicegroupmember_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("servicegroup_servicegroupmember_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountServicegroupServicegroupmemberBinding(id string) (int, error) {
	var results count_servicegroup_servicegroupmember_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("servicegroup_servicegroupmember_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsServicegroupServicegroupmemberBinding(key ServicegroupServicegroupmemberBindingKey) (bool, error) {
	var results count_servicegroup_servicegroupmember_binding_result

	id, qs := servicegroup_servicegroupmember_binding_key_to_id_filter(key)

	qs["count"] = "yes"

	if err := c.get("servicegroup_servicegroupmember_binding", id, qs, &results); err != nil {
		return false, err
	} else {
		if len(results.Results) > 1 {
			return false, fmt.Errorf("More than one servicegroup_servicegroupmember_binding element found")
		}

		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) BulkListServicegroupServicegroupmemberBinding() ([]ServicegroupServicegroupmemberBinding, error) {
	var results get_servicegroup_servicegroupmember_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("servicegroup_servicegroupmember_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListServicegroupServicegroupmemberBinding(id string) ([]ServicegroupServicegroupmemberBinding, error) {
	var results get_servicegroup_servicegroupmember_binding_result

	if err := c.get("servicegroup_servicegroupmember_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetServicegroupServicegroupmemberBinding(key ServicegroupServicegroupmemberBindingKey) (*ServicegroupServicegroupmemberBinding, error) {
	var results get_servicegroup_servicegroupmember_binding_result

	id, qs := servicegroup_servicegroupmember_binding_key_to_id_args(key)

	if err := c.get("servicegroup_servicegroupmember_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one servicegroup_servicegroupmember_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("servicegroup_servicegroupmember_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteServicegroupServicegroupmemberBinding(key ServicegroupServicegroupmemberBindingKey) error {
	id, qs := servicegroup_servicegroupmember_binding_key_to_id_args(key)

	return c.delete("servicegroup_servicegroupmember_binding", id, qs)
}
