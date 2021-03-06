package nitro

import (
	"fmt"
	"net/url"
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

func (key ServicegroupLbmonitorBindingKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = url.QueryEscape(key.Servicegroupname)
	args = append(args, "monitor_name:"+url.QueryEscape(key.Monitor_name))

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key ServicegroupLbmonitorBindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key ServicegroupLbmonitorBindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
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

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("servicegroup_lbmonitor_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_servicegroup_lbmonitor_binding_result struct {
	Results []Count `json:"servicegroup_lbmonitor_binding"`
}

func (c *NitroClient) CountServicegroupLbmonitorBinding() (int, error) {
	results := count_servicegroup_lbmonitor_binding_result{}

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

//      READ

type get_servicegroup_lbmonitor_binding_result struct {
	Results []ServicegroupLbmonitorBinding `json:"servicegroup_lbmonitor_binding"`
}

func (c *NitroClient) GetServicegroupLbmonitorBinding(key ServicegroupLbmonitorBindingKey) (*ServicegroupLbmonitorBinding, error) {
	var results get_servicegroup_lbmonitor_binding_result

	id, qs := key.to_id_filter()

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

func (c *NitroClient) ExistsServicegroupLbmonitorBinding(key ServicegroupLbmonitorBindingKey) (bool, error) {
	var results count_servicegroup_lbmonitor_binding_result

	id, qs := key.to_id_filter()

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
