package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Tmtrafficaction struct {
	Apptimeout       int    `json:"apptimeout,string,omitempty"`
	Forcedtimeout    string `json:"forcedtimeout,omitempty"`
	Forcedtimeoutval int    `json:"forcedtimeoutval,string,omitempty"`
	Formssoaction    string `json:"formssoaction,omitempty"`
	Initiatelogout   string `json:"initiatelogout,omitempty"`
	Kcdaccount       string `json:"kcdaccount,omitempty"`
	Name             string `json:"name,omitempty"`
	Passwdexpression string `json:"passwdexpression,omitempty"`
	Persistentcookie string `json:"persistentcookie,omitempty"`
	Samlssoprofile   string `json:"samlssoprofile,omitempty"`
	Sso              string `json:"sso,omitempty"`
	Userexpression   string `json:"userexpression,omitempty"`
}

type TmtrafficactionKey struct {
	Name string
}

func (resource Tmtrafficaction) ToKey() TmtrafficactionKey {
	key := TmtrafficactionKey{
		resource.Name,
	}

	return key
}

func (key TmtrafficactionKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key TmtrafficactionKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key TmtrafficactionKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_tmtrafficaction_payload struct {
	Resource Tmtrafficaction `json:"tmtrafficaction"`
}

func (c *NitroClient) AddTmtrafficaction(resource Tmtrafficaction) error {
	payload := add_tmtrafficaction_payload{
		resource,
	}

	return c.post("tmtrafficaction", "", nil, payload)
}

//      LIST

type list_tmtrafficaction_result struct {
	Results []Tmtrafficaction `json:"tmtrafficaction"`
}

func (c *NitroClient) ListTmtrafficaction() ([]Tmtrafficaction, error) {
	results := list_tmtrafficaction_result{}

	if err := c.get("tmtrafficaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_tmtrafficaction_result struct {
	Results []Tmtrafficaction `json:"tmtrafficaction"`
}

func (c *NitroClient) GetTmtrafficaction(key TmtrafficactionKey) (*Tmtrafficaction, error) {
	var results get_tmtrafficaction_result

	id, qs := key.to_id_args()

	if err := c.get("tmtrafficaction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one tmtrafficaction element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("tmtrafficaction element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_tmtrafficaction_result struct {
	Results []Count `json:"tmtrafficaction"`
}

func (c *NitroClient) CountTmtrafficaction() (int, error) {
	var results count_tmtrafficaction_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("tmtrafficaction", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsTmtrafficaction(key TmtrafficactionKey) (bool, error) {
	var results count_tmtrafficaction_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("tmtrafficaction", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteTmtrafficaction(key TmtrafficactionKey) error {
	id, qs := key.to_id_args()

	return c.delete("tmtrafficaction", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
