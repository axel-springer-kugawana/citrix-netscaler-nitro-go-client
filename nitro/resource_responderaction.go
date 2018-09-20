package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Responderaction struct {
	Bypasssafetycheck  string `json:"bypasssafetycheck,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Htmlpage           string `json:"htmlpage,omitempty"`
	Name               string `json:"name,omitempty"`
	Reasonphrase       string `json:"reasonphrase,omitempty"`
	Responsestatuscode int    `json:"responsestatuscode,string,omitempty"`
	Target             string `json:"target,omitempty"`
	Type               string `json:"type,omitempty"`
}

type ResponderactionKey struct {
	Name string
}

func (resource Responderaction) ToKey() ResponderactionKey {
	key := ResponderactionKey{
		resource.Name,
	}

	return key
}

func (key ResponderactionKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = url.QueryEscape(key.Name)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key ResponderactionKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key ResponderactionKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_responderaction_payload struct {
	Resource Responderaction `json:"responderaction"`
}

func (c *NitroClient) AddResponderaction(resource Responderaction) error {
	payload := add_responderaction_payload{
		resource,
	}

	return c.post("responderaction", "", nil, payload)
}

//      LIST

type list_responderaction_result struct {
	Results []Responderaction `json:"responderaction"`
}

func (c *NitroClient) ListResponderaction() ([]Responderaction, error) {
	results := list_responderaction_result{}

	if err := c.get("responderaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_responderaction_result struct {
	Results []Responderaction `json:"responderaction"`
}

func (c *NitroClient) GetResponderaction(key ResponderactionKey) (*Responderaction, error) {
	var results get_responderaction_result

	id, qs := key.to_id_args()

	if err := c.get("responderaction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one responderaction element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("responderaction element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_responderaction_result struct {
	Results []Count `json:"responderaction"`
}

func (c *NitroClient) CountResponderaction() (int, error) {
	var results count_responderaction_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("responderaction", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsResponderaction(key ResponderactionKey) (bool, error) {
	var results count_responderaction_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("responderaction", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteResponderaction(key ResponderactionKey) error {
	id, qs := key.to_id_args()

	return c.delete("responderaction", id, qs)
}

//      UPDATE

type ResponderactionUpdate struct {
	Name               string `json:"name,omitempty"`
	Target             string `json:"target,omitempty"`
	Bypasssafetycheck  string `json:"bypasssafetycheck,omitempty"`
	Htmlpage           string `json:"htmlpage,omitempty"`
	Responsestatuscode int    `json:"responsestatuscode,string,omitempty"`
	Reasonphrase       string `json:"reasonphrase,omitempty"`
	Comment            string `json:"comment,omitempty"`
}

func (resource Responderaction) ToUpdate() ResponderactionUpdate {
	update := ResponderactionUpdate{
		resource.Name,
		resource.Target,
		resource.Bypasssafetycheck,
		resource.Htmlpage,
		resource.Responsestatuscode,
		resource.Reasonphrase,
		resource.Comment,
	}

	return update
}

type update_responderaction_payload struct {
	Update ResponderactionUpdate `json:"responderaction"`
}

func (c *NitroClient) UpdateResponderaction(update ResponderactionUpdate) error {
	payload := update_responderaction_payload{
		update,
	}

	return c.put("responderaction", "", nil, payload)
}

//      UNSET

type ResponderactionUnset struct {
	Name               string `json:"name,omitempty"`
	Target             bool   `json:"target,omitempty"`
	Bypasssafetycheck  bool   `json:"bypasssafetycheck,omitempty"`
	Htmlpage           bool   `json:"htmlpage,omitempty"`
	Responsestatuscode bool   `json:"responsestatuscode,omitempty"`
	Reasonphrase       bool   `json:"reasonphrase,omitempty"`
	Comment            bool   `json:"comment,omitempty"`
}

func (resource Responderaction) ToUnset() ResponderactionUnset {
	unset := ResponderactionUnset{
		resource.Name,
		false,
		false,
		false,
		false,
		false,
		false,
	}

	return unset
}

type unset_responderaction_payload struct {
	Unset ResponderactionUnset `json:"responderaction"`
}

func (c *NitroClient) UnsetResponderaction(unset ResponderactionUnset) error {
	payload := unset_responderaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("responderaction", "", qs, payload)
}

//      RENAME
//      TODO
