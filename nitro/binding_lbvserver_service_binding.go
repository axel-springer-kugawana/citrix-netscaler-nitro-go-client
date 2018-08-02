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

type get_lbvserver_service_binding struct {
	Results []LbvserverServiceBinding `json:"lbvserver_service_binding"`
}

type add_lbvserver_service_binding_payload struct {
	lbvserver_service_binding LbvserverServiceBinding
}

func lbvserver_service_binding_key_to_id_args(key LbvserverServiceBindingKey) (string, string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)
	args = append(args, "servicename:"+key.Servicename)

	return "", strings.Join(args, ",")
}

// TODO : Exists
// TODO : Count

func (c *NitroClient) AddLbvserverServiceBinding(binding LbvserverServiceBinding) error {
	payload := add_lbvserver_service_binding_payload{
		binding,
	}

	return c.put("lbvserver_service_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbvserverServiceBinding(key LbvserverServiceBindingKey) ([]LbvserverServiceBinding, error) {
	var results get_lbvserver_service_binding

	id, args := lbvserver_service_binding_key_to_id_args(key)

	if err := c.get("lbvserver_service_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbvserverServiceBinding() ([]LbvserverServiceBinding, error) {
	var results get_lbvserver_service_binding

	if err := c.get("lbvserver_service_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverServiceBinding(key LbvserverServiceBindingKey) (*LbvserverServiceBinding, error) {
	var results get_lbvserver_service_binding

	id, args := lbvserver_service_binding_key_to_id_args(key)

	if err := c.get("lbvserver_service_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_service_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("lbvserver_service_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverServiceBinding(key LbvserverServiceBindingKey) error {
	id, args := lbvserver_service_binding_key_to_id_args(key)

	return c.delete("lbvserver_service_binding", id, "", args)
}
