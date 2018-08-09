package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Scpolicy struct {
	Action            string `json:"action,omitempty"`
	Altcontentpath    string `json:"altcontentpath,omitempty"`
	Altcontentsvcname string `json:"altcontentsvcname,omitempty"`
	Delay             int    `json:"delay,string,omitempty"`
	Maxconn           int    `json:"maxconn,string,omitempty"`
	Name              string `json:"name,omitempty"`
	Rule              string `json:"rule,omitempty"`
	Url               string `json:"url,omitempty"`
}

type ScpolicyKey struct {
	Name string
}

func (resource Scpolicy) ToKey() ScpolicyKey {
	key := ScpolicyKey{
		resource.Name,
	}

	return key
}

func (key ScpolicyKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Name

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key ScpolicyKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key ScpolicyKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_scpolicy_payload struct {
	Resource Scpolicy `json:"scpolicy"`
}

func (c *NitroClient) AddScpolicy(resource Scpolicy) error {
	payload := add_scpolicy_payload{
		resource,
	}

	return c.post("scpolicy", "", nil, payload)
}

//      LIST

type list_scpolicy_result struct {
	Results []Scpolicy `json:"scpolicy"`
}

func (c *NitroClient) ListScpolicy() ([]Scpolicy, error) {
	results := list_scpolicy_result{}

	if err := c.get("scpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_scpolicy_result struct {
	Results []Scpolicy `json:"scpolicy"`
}

func (c *NitroClient) GetScpolicy(key ScpolicyKey) (*Scpolicy, error) {
	var results get_scpolicy_result

	id, qs := key.to_id_args()

	if err := c.get("scpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one scpolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("scpolicy element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_scpolicy_result struct {
	Results []Count `json:"scpolicy"`
}

func (c *NitroClient) CountScpolicy() (int, error) {
	var results count_scpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("scpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsScpolicy(key ScpolicyKey) (bool, error) {
	var results count_scpolicy_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("scpolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteScpolicy(key ScpolicyKey) error {
	id, qs := key.to_id_args()

	return c.delete("scpolicy", id, qs)
}

//      UPDATE

type ScpolicyUpdate struct {
	Name              string `json:"name,omitempty"`
	Url               string `json:"url,omitempty"`
	Rule              string `json:"rule,omitempty"`
	Delay             int    `json:"delay,string,omitempty"`
	Maxconn           int    `json:"maxconn,string,omitempty"`
	Action            string `json:"action,omitempty"`
	Altcontentsvcname string `json:"altcontentsvcname,omitempty"`
	Altcontentpath    string `json:"altcontentpath,omitempty"`
}

func (resource Scpolicy) ToUpdate() ScpolicyUpdate {
	update := ScpolicyUpdate{
		resource.Name,
		resource.Url,
		resource.Rule,
		resource.Delay,
		resource.Maxconn,
		resource.Action,
		resource.Altcontentsvcname,
		resource.Altcontentpath,
	}

	return update
}

type update_scpolicy_payload struct {
	Update ScpolicyUpdate `json:"scpolicy"`
}

func (c *NitroClient) UpdateScpolicy(update ScpolicyUpdate) error {
	payload := update_scpolicy_payload{
		update,
	}

	return c.put("scpolicy", "", nil, payload)
}

//      UNSET

type ScpolicyUnset struct {
	Name              string `json:"name,omitempty"`
	Url               bool   `json:"url,omitempty"`
	Rule              bool   `json:"rule,omitempty"`
	Delay             bool   `json:"delay,omitempty"`
	Maxconn           bool   `json:"maxconn,omitempty"`
	Action            bool   `json:"action,omitempty"`
	Altcontentsvcname bool   `json:"altcontentsvcname,omitempty"`
	Altcontentpath    bool   `json:"altcontentpath,omitempty"`
}

func (resource Scpolicy) ToUnset() ScpolicyUnset {
	unset := ScpolicyUnset{
		resource.Name,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
	}

	return unset
}

type unset_scpolicy_payload struct {
	Unset ScpolicyUnset `json:"scpolicy"`
}

func (c *NitroClient) UnsetScpolicy(unset ScpolicyUnset) error {
	payload := unset_scpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("scpolicy", "", qs, payload)
}

//      RENAME
//      TODO
