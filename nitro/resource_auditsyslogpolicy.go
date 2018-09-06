package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Auditsyslogpolicy struct {
	Action string `json:"action,omitempty"`
	Name   string `json:"name,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

type AuditsyslogpolicyKey struct {
	Name string
}

func (resource Auditsyslogpolicy) ToKey() AuditsyslogpolicyKey {
	key := AuditsyslogpolicyKey{
		resource.Name,
	}

	return key
}

func (key AuditsyslogpolicyKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = key.Name

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key AuditsyslogpolicyKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key AuditsyslogpolicyKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_auditsyslogpolicy_payload struct {
	Resource Auditsyslogpolicy `json:"auditsyslogpolicy"`
}

func (c *NitroClient) AddAuditsyslogpolicy(resource Auditsyslogpolicy) error {
	payload := add_auditsyslogpolicy_payload{
		resource,
	}

	return c.post("auditsyslogpolicy", "", nil, payload)
}

//      LIST

type list_auditsyslogpolicy_result struct {
	Results []Auditsyslogpolicy `json:"auditsyslogpolicy"`
}

func (c *NitroClient) ListAuditsyslogpolicy() ([]Auditsyslogpolicy, error) {
	results := list_auditsyslogpolicy_result{}

	if err := c.get("auditsyslogpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_auditsyslogpolicy_result struct {
	Results []Auditsyslogpolicy `json:"auditsyslogpolicy"`
}

func (c *NitroClient) GetAuditsyslogpolicy(key AuditsyslogpolicyKey) (*Auditsyslogpolicy, error) {
	var results get_auditsyslogpolicy_result

	id, qs := key.to_id_args()

	if err := c.get("auditsyslogpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one auditsyslogpolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("auditsyslogpolicy element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_auditsyslogpolicy_result struct {
	Results []Count `json:"auditsyslogpolicy"`
}

func (c *NitroClient) CountAuditsyslogpolicy() (int, error) {
	var results count_auditsyslogpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("auditsyslogpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsAuditsyslogpolicy(key AuditsyslogpolicyKey) (bool, error) {
	var results count_auditsyslogpolicy_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("auditsyslogpolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteAuditsyslogpolicy(key AuditsyslogpolicyKey) error {
	id, qs := key.to_id_args()

	return c.delete("auditsyslogpolicy", id, qs)
}

//      UPDATE

type AuditsyslogpolicyUpdate struct {
	Name   string `json:"name,omitempty"`
	Rule   string `json:"rule,omitempty"`
	Action string `json:"action,omitempty"`
}

func (resource Auditsyslogpolicy) ToUpdate() AuditsyslogpolicyUpdate {
	update := AuditsyslogpolicyUpdate{
		resource.Name,
		resource.Rule,
		resource.Action,
	}

	return update
}

type update_auditsyslogpolicy_payload struct {
	Update AuditsyslogpolicyUpdate `json:"auditsyslogpolicy"`
}

func (c *NitroClient) UpdateAuditsyslogpolicy(update AuditsyslogpolicyUpdate) error {
	payload := update_auditsyslogpolicy_payload{
		update,
	}

	return c.put("auditsyslogpolicy", "", nil, payload)
}

//      UNSET

type AuditsyslogpolicyUnset struct {
	Name   string `json:"name,omitempty"`
	Rule   bool   `json:"rule,omitempty"`
	Action bool   `json:"action,omitempty"`
}

func (resource Auditsyslogpolicy) ToUnset() AuditsyslogpolicyUnset {
	unset := AuditsyslogpolicyUnset{
		resource.Name,
		false,
		false,
	}

	return unset
}

type unset_auditsyslogpolicy_payload struct {
	Unset AuditsyslogpolicyUnset `json:"auditsyslogpolicy"`
}

func (c *NitroClient) UnsetAuditsyslogpolicy(unset AuditsyslogpolicyUnset) error {
	payload := unset_auditsyslogpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("auditsyslogpolicy", "", qs, payload)
}

//      RENAME
//      TODO
