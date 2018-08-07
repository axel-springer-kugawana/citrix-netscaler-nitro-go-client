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

func (resource ServicegroupLbmonitorBinding) ToKey() ServicegroupLbmonitorBindingKey {
	key := ServicegroupLbmonitorBindingKey{
		resource.Servicegroupname,
		resource.Monitor_name,
	}

	return key
}

func (key ServicegroupLbmonitorBindingKey) to_id_args() (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Servicegroupname
	args = append(args, "monitor_name:"+key.Monitor_name)

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return id, qs
}

//      CREATE

type add_servicegroup_lbmonitor_binding_payload struct {
	Resource ServicegroupLbmonitorBinding `json:"servicegroup_lbmonitor_binding"`
}

func (c *NitroClient) AddServicegroupLbmonitorBinding(resource ServicegroupLbmonitorBinding) error {
	payload := add_servicegroup_lbmonitor_binding_payload{
		resource,
	}

	return c.put("servicegroup_lbmonitor_binding", "", nil, payload)
}

//      LIST

type list_servicegroup_lbmonitor_binding_result struct {
	Results []ServicegroupLbmonitorBinding `json:"servicegroup_lbmonitor_binding"`
}

func (c *NitroClient) ListServicegroupLbmonitorBinding() ([]ServicegroupLbmonitorBinding, error) {
	results := list_servicegroup_lbmonitor_binding_result{}

	if err := c.get("servicegroup_lbmonitor_binding", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_servicegroup_lbmonitor_binding_result struct {
	Results []ServicegroupLbmonitorBinding `json:"servicegroup_lbmonitor_binding"`
}

func (c *NitroClient) GetServicegroupLbmonitorBinding(key ServicegroupLbmonitorBindingKey) (*ServicegroupLbmonitorBinding, error) {
	var results get_servicegroup_lbmonitor_binding_result

	id, qs := key.to_id_args()

	if err := c.get("servicegroup_lbmonitor_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one servicegroup_lbmonitor_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("servicegroup_lbmonitor_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

type count_servicegroup_lbmonitor_binding_result struct {
	Results []Count `json:"servicegroup_lbmonitor_binding"`
}

func (c *NitroClient) ExistsServicegroupLbmonitorBinding(key ServicegroupLbmonitorBindingKey) (bool, error) {
	var results count_servicegroup_lbmonitor_binding_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("servicegroup_lbmonitor_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteServicegroupLbmonitorBinding(key ServicegroupLbmonitorBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("servicegroup_lbmonitor_binding", id, qs)
}
