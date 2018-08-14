package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type CsvserverFeopolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type CsvserverFeopolicyBindingKey struct {
	Name       string
	Policyname string
}

func (resource CsvserverFeopolicyBinding) ToKey() CsvserverFeopolicyBindingKey {
	key := CsvserverFeopolicyBindingKey{
		resource.Name,
		resource.Policyname,
	}

	return key
}

func (key CsvserverFeopolicyBindingKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key CsvserverFeopolicyBindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key CsvserverFeopolicyBindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_csvserver_feopolicy_binding_payload struct {
	Resource CsvserverFeopolicyBinding `json:"csvserver_feopolicy_binding"`
}

func (c *NitroClient) AddCsvserverFeopolicyBinding(resource CsvserverFeopolicyBinding) error {
	payload := add_csvserver_feopolicy_binding_payload{
		resource,
	}

	return c.put("csvserver_feopolicy_binding", "", nil, payload)
}

//      LIST

type list_csvserver_feopolicy_binding_result struct {
	Results []CsvserverFeopolicyBinding `json:"csvserver_feopolicy_binding"`
}

func (c *NitroClient) ListCsvserverFeopolicyBinding() ([]CsvserverFeopolicyBinding, error) {
	results := list_csvserver_feopolicy_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("csvserver_feopolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_csvserver_feopolicy_binding_result struct {
	Results []Count `json:"csvserver_feopolicy_binding"`
}

func (c *NitroClient) CountCsvserverFeopolicyBinding() (int, error) {
	results := count_csvserver_feopolicy_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("csvserver_feopolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      READ

type get_csvserver_feopolicy_binding_result struct {
	Results []CsvserverFeopolicyBinding `json:"csvserver_feopolicy_binding"`
}

func (c *NitroClient) GetCsvserverFeopolicyBinding(key CsvserverFeopolicyBindingKey) (*CsvserverFeopolicyBinding, error) {
	var results get_csvserver_feopolicy_binding_result

	id, qs := key.to_id_filter()

	if err := c.get("csvserver_feopolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one csvserver_feopolicy_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("csvserver_feopolicy_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

func (c *NitroClient) ExistsCsvserverFeopolicyBinding(key CsvserverFeopolicyBindingKey) (bool, error) {
	var results count_csvserver_feopolicy_binding_result

	id, qs := key.to_id_filter()

	qs["count"] = "yes"

	if err := c.get("csvserver_feopolicy_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteCsvserverFeopolicyBinding(key CsvserverFeopolicyBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("csvserver_feopolicy_binding", id, qs)
}
