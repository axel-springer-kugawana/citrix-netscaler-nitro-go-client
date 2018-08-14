package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type CsvserverTmtrafficpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type CsvserverTmtrafficpolicyBindingKey struct {
	Name       string
	Policyname string
}

func (resource CsvserverTmtrafficpolicyBinding) ToKey() CsvserverTmtrafficpolicyBindingKey {
	key := CsvserverTmtrafficpolicyBindingKey{
		resource.Name,
		resource.Policyname,
	}

	return key
}

func (key CsvserverTmtrafficpolicyBindingKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Name
	args = append(args, "policyname:"+key.Policyname)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key CsvserverTmtrafficpolicyBindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key CsvserverTmtrafficpolicyBindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_csvserver_tmtrafficpolicy_binding_payload struct {
	Resource CsvserverTmtrafficpolicyBinding `json:"csvserver_tmtrafficpolicy_binding"`
}

func (c *NitroClient) AddCsvserverTmtrafficpolicyBinding(resource CsvserverTmtrafficpolicyBinding) error {
	payload := add_csvserver_tmtrafficpolicy_binding_payload{
		resource,
	}

	return c.put("csvserver_tmtrafficpolicy_binding", "", nil, payload)
}

//      LIST

type list_csvserver_tmtrafficpolicy_binding_result struct {
	Results []CsvserverTmtrafficpolicyBinding `json:"csvserver_tmtrafficpolicy_binding"`
}

func (c *NitroClient) ListCsvserverTmtrafficpolicyBinding() ([]CsvserverTmtrafficpolicyBinding, error) {
	results := list_csvserver_tmtrafficpolicy_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("csvserver_tmtrafficpolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_csvserver_tmtrafficpolicy_binding_result struct {
	Results []Count `json:"csvserver_tmtrafficpolicy_binding"`
}

func (c *NitroClient) CountCsvserverTmtrafficpolicyBinding() (int, error) {
	results := count_csvserver_tmtrafficpolicy_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("csvserver_tmtrafficpolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      READ

type get_csvserver_tmtrafficpolicy_binding_result struct {
	Results []CsvserverTmtrafficpolicyBinding `json:"csvserver_tmtrafficpolicy_binding"`
}

func (c *NitroClient) GetCsvserverTmtrafficpolicyBinding(key CsvserverTmtrafficpolicyBindingKey) (*CsvserverTmtrafficpolicyBinding, error) {
	var results get_csvserver_tmtrafficpolicy_binding_result

	id, qs := key.to_id_filter()

	if err := c.get("csvserver_tmtrafficpolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one csvserver_tmtrafficpolicy_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("csvserver_tmtrafficpolicy_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

func (c *NitroClient) ExistsCsvserverTmtrafficpolicyBinding(key CsvserverTmtrafficpolicyBindingKey) (bool, error) {
	var results count_csvserver_tmtrafficpolicy_binding_result

	id, qs := key.to_id_filter()

	qs["count"] = "yes"

	if err := c.get("csvserver_tmtrafficpolicy_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteCsvserverTmtrafficpolicyBinding(key CsvserverTmtrafficpolicyBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("csvserver_tmtrafficpolicy_binding", id, qs)
}
