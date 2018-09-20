package nitro

import (
	"fmt"
	"net/url"
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

func (resource ServicegroupServicegroupmemberBinding) ToKey() ServicegroupServicegroupmemberBindingKey {
	key := ServicegroupServicegroupmemberBindingKey{
		resource.Servicegroupname,
		resource.Servername,
		resource.Port,
	}

	return key
}

func (key ServicegroupServicegroupmemberBindingKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = url.QueryEscape(key.Servicegroupname)
	args = append(args, "servername:"+url.QueryEscape(key.Servername))
	args = append(args, "port:"+strconv.Itoa(key.Port))

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key ServicegroupServicegroupmemberBindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key ServicegroupServicegroupmemberBindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_servicegroup_servicegroupmember_binding_payload struct {
	Resource ServicegroupServicegroupmemberBinding `json:"servicegroup_servicegroupmember_binding"`
}

func (c *NitroClient) AddServicegroupServicegroupmemberBinding(resource ServicegroupServicegroupmemberBinding) error {
	payload := add_servicegroup_servicegroupmember_binding_payload{
		resource,
	}

	return c.put("servicegroup_servicegroupmember_binding", "", nil, payload)
}

//      LIST

type list_servicegroup_servicegroupmember_binding_result struct {
	Results []ServicegroupServicegroupmemberBinding `json:"servicegroup_servicegroupmember_binding"`
}

func (c *NitroClient) ListServicegroupServicegroupmemberBinding() ([]ServicegroupServicegroupmemberBinding, error) {
	results := list_servicegroup_servicegroupmember_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("servicegroup_servicegroupmember_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_servicegroup_servicegroupmember_binding_result struct {
	Results []Count `json:"servicegroup_servicegroupmember_binding"`
}

func (c *NitroClient) CountServicegroupServicegroupmemberBinding() (int, error) {
	results := count_servicegroup_servicegroupmember_binding_result{}

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

//      READ

type get_servicegroup_servicegroupmember_binding_result struct {
	Results []ServicegroupServicegroupmemberBinding `json:"servicegroup_servicegroupmember_binding"`
}

func (c *NitroClient) GetServicegroupServicegroupmemberBinding(key ServicegroupServicegroupmemberBindingKey) (*ServicegroupServicegroupmemberBinding, error) {
	var results get_servicegroup_servicegroupmember_binding_result

	id, qs := key.to_id_filter()

	if err := c.get("servicegroup_servicegroupmember_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one servicegroup_servicegroupmember_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("servicegroup_servicegroupmember_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

func (c *NitroClient) ExistsServicegroupServicegroupmemberBinding(key ServicegroupServicegroupmemberBindingKey) (bool, error) {
	var results count_servicegroup_servicegroupmember_binding_result

	id, qs := key.to_id_filter()

	qs["count"] = "yes"

	if err := c.get("servicegroup_servicegroupmember_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteServicegroupServicegroupmemberBinding(key ServicegroupServicegroupmemberBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("servicegroup_servicegroupmember_binding", id, qs)
}
