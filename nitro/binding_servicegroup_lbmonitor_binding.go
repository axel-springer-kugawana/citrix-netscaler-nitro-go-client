package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type ServicegroupLbmonitorBinding struct {
	Monitor_name     string `json:"monitor_name,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Weight           int    `json:"weight,string,omitempty"`
}

type ServicegroupLbmonitorBindingKey struct {
	Servicegroupname string
	Monitor_name     string
}

type add_servicegroup_lbmonitor_binding_payload struct {
	Resource ServicegroupLbmonitorBinding `json:"servicegroup_lbmonitor_binding"`
}

type get_servicegroup_lbmonitor_binding_result struct {
	Results []ServicegroupLbmonitorBinding `json:"servicegroup_lbmonitor_binding"`
}

type count_servicegroup_lbmonitor_binding_result struct {
	Results []Count `json:"servicegroup_lbmonitor_binding"`
}

func servicegroup_lbmonitor_binding_key_to_id_args(key ServicegroupLbmonitorBindingKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "servicegroupname:"+key.Servicegroupname)
	args = append(args, "monitor_name:"+key.Monitor_name)

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return "", qs
}

func (c *NitroClient) AddServicegroupLbmonitorBinding(binding ServicegroupLbmonitorBinding) error {
	payload := add_servicegroup_lbmonitor_binding_payload{
		binding,
	}

	return c.put("servicegroup_lbmonitor_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountServicegroupLbmonitorBinding() (int, error) {
	var results count_servicegroup_lbmonitor_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("servicegroup_lbmonitor_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountServicegroupLbmonitorBinding(id string) (int, error) {
	var results count_servicegroup_lbmonitor_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("servicegroup_lbmonitor_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsServicegroupLbmonitorBinding(key ServicegroupLbmonitorBindingKey) (bool, error) {
	// TODO : wrong implementation
	return false, nil
	//        if count, err := c.CountServicegroupLbmonitorBinding(id); err != nil {
	//                return false, err
	//        } else {
	//                return count == 1, nil
	//        }
}

func (c *NitroClient) BulkListServicegroupLbmonitorBinding() ([]ServicegroupLbmonitorBinding, error) {
	var results get_servicegroup_lbmonitor_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("servicegroup_lbmonitor_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListServicegroupLbmonitorBinding(id string) ([]ServicegroupLbmonitorBinding, error) {
	var results get_servicegroup_lbmonitor_binding_result

	if err := c.get("servicegroup_lbmonitor_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetServicegroupLbmonitorBinding(key ServicegroupLbmonitorBindingKey) (*ServicegroupLbmonitorBinding, error) {
	var results get_servicegroup_lbmonitor_binding_result

	id, qs := servicegroup_lbmonitor_binding_key_to_id_args(key)

	if err := c.get("servicegroup_lbmonitor_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one servicegroup_lbmonitor_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("servicegroup_lbmonitor_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteServicegroupLbmonitorBinding(key ServicegroupLbmonitorBindingKey) error {
	id, qs := servicegroup_lbmonitor_binding_key_to_id_args(key)

	return c.delete("servicegroup_lbmonitor_binding", id, qs)
}
