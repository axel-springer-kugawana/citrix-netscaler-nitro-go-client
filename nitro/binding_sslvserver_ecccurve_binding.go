package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type SslvserverEcccurveBinding struct {
	Ecccurvename string `json:"ecccurvename,omitempty"`
	Vservername  string `json:"vservername,omitempty"`
}

type SslvserverEcccurveBindingKey struct {
	Vservername  string
	Ecccurvename string
}

func (resource SslvserverEcccurveBinding) ToKey() SslvserverEcccurveBindingKey {
	key := SslvserverEcccurveBindingKey{
		resource.Vservername,
		resource.Ecccurvename,
	}

	return key
}

func (key SslvserverEcccurveBindingKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = key.Vservername
	args = append(args, "ecccurvename:"+key.Ecccurvename)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key SslvserverEcccurveBindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key SslvserverEcccurveBindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_sslvserver_ecccurve_binding_payload struct {
	Resource SslvserverEcccurveBinding `json:"sslvserver_ecccurve_binding"`
}

func (c *NitroClient) AddSslvserverEcccurveBinding(resource SslvserverEcccurveBinding) error {
	payload := add_sslvserver_ecccurve_binding_payload{
		resource,
	}

	return c.put("sslvserver_ecccurve_binding", "", nil, payload)
}

//      LIST

type list_sslvserver_ecccurve_binding_result struct {
	Results []SslvserverEcccurveBinding `json:"sslvserver_ecccurve_binding"`
}

func (c *NitroClient) ListSslvserverEcccurveBinding() ([]SslvserverEcccurveBinding, error) {
	results := list_sslvserver_ecccurve_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("sslvserver_ecccurve_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_sslvserver_ecccurve_binding_result struct {
	Results []Count `json:"sslvserver_ecccurve_binding"`
}

func (c *NitroClient) CountSslvserverEcccurveBinding() (int, error) {
	results := count_sslvserver_ecccurve_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("sslvserver_ecccurve_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      READ

type get_sslvserver_ecccurve_binding_result struct {
	Results []SslvserverEcccurveBinding `json:"sslvserver_ecccurve_binding"`
}

func (c *NitroClient) GetSslvserverEcccurveBinding(key SslvserverEcccurveBindingKey) (*SslvserverEcccurveBinding, error) {
	var results get_sslvserver_ecccurve_binding_result

	id, qs := key.to_id_filter()

	if err := c.get("sslvserver_ecccurve_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one sslvserver_ecccurve_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("sslvserver_ecccurve_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

func (c *NitroClient) ExistsSslvserverEcccurveBinding(key SslvserverEcccurveBindingKey) (bool, error) {
	var results count_sslvserver_ecccurve_binding_result

	id, qs := key.to_id_filter()

	qs["count"] = "yes"

	if err := c.get("sslvserver_ecccurve_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteSslvserverEcccurveBinding(key SslvserverEcccurveBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("sslvserver_ecccurve_binding", id, qs)
}
