package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Responderaction struct {
	Name               string `json:"name"`
	Bypasssafetycheck  string `json:"bypasssafetycheck,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Htmlpage           string `json:"htmlpage,omitempty"`
	Reasonphrase       string `json:"reasonphrase,omitempty"`
	Responsestatuscode int    `json:"responsestatuscode,string,omitempty"`
	Target             string `json:"target,omitempty"`
	Type               string `json:"type,omitempty"`
}

func responderaction_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type ResponderactionUnset struct {
	Name               string `json:"name"`
	Target             bool   `json:"target,string,omitempty"`
	Bypasssafetycheck  bool   `json:"bypasssafetycheck,string,omitempty"`
	Htmlpage           bool   `json:"htmlpage,string,omitempty"`
	Responsestatuscode bool   `json:"responsestatuscode,string,omitempty"`
	Reasonphrase       bool   `json:"reasonphrase,string,omitempty"`
	Comment            bool   `json:"comment,string,omitempty"`
}

type update_responderaction struct {
	Name               string `json:"name"`
	Target             string `json:"target,omitempty"`
	Bypasssafetycheck  string `json:"bypasssafetycheck,omitempty"`
	Htmlpage           string `json:"htmlpage,omitempty"`
	Responsestatuscode int    `json:"responsestatuscode,string,omitempty"`
	Reasonphrase       string `json:"reasonphrase,omitempty"`
	Comment            string `json:"comment,omitempty"`
}

type rename_responderaction struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_responderaction_payload struct {
	Resource Responderaction `json:"responderaction"`
}

type rename_responderaction_payload struct {
	Rename rename_responderaction `json:"responderaction"`
}

type unset_responderaction_payload struct {
	Unset ResponderactionUnset `json:"responderaction"`
}

type update_responderaction_payload struct {
	Update update_responderaction `json:"responderaction"`
}

type get_responderaction_result struct {
	Results []Responderaction `json:"responderaction"`
}

type count_responderaction_result struct {
	Results []Count `json:"responderaction"`
}

func (c *NitroClient) AddResponderaction(resource Responderaction) error {
	payload := add_responderaction_payload{
		resource,
	}

	return c.post("responderaction", "", nil, payload)
}

func (c *NitroClient) RenameResponderaction(name string, newName string) error {
	payload := rename_responderaction_payload{
		rename_responderaction{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("responderaction", "", qs, payload)
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

func (c *NitroClient) ExistsResponderaction(key string) (bool, error) {
	var results count_responderaction_result

	id, qs := responderaction_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("responderaction", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListResponderaction() ([]Responderaction, error) {
	results := get_responderaction_result{}

	if err := c.get("responderaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetResponderaction(key string) (*Responderaction, error) {
	var results get_responderaction_result

	id, qs := responderaction_key_to_id_args(key)

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

func (c *NitroClient) DeleteResponderaction(key string) error {
	id, qs := responderaction_key_to_id_args(key)

	return c.delete("responderaction", id, qs)
}

func (c *NitroClient) UnsetResponderaction(unset ResponderactionUnset) error {
	payload := unset_responderaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("responderaction", "", qs, payload)
}

func (c *NitroClient) UpdateResponderaction(resource Responderaction) error {
	payload := update_responderaction_payload{
		update_responderaction{
			resource.Name,
			resource.Target,
			resource.Bypasssafetycheck,
			resource.Htmlpage,
			resource.Responsestatuscode,
			resource.Reasonphrase,
			resource.Comment,
		},
	}

	return c.put("responderaction", "", nil, payload)
}
