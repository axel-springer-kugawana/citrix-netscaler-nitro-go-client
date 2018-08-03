package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Scpolicy struct {
	Name              string `json:"name"`
	Action            string `json:"action,omitempty"`
	Altcontentpath    string `json:"altcontentpath,omitempty"`
	Altcontentsvcname string `json:"altcontentsvcname,omitempty"`
	Delay             int    `json:"delay,string,omitempty"`
	Maxconn           int    `json:"maxconn,string,omitempty"`
	Rule              string `json:"rule,omitempty"`
	Url               string `json:"url,omitempty"`
}

func scpolicy_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type ScpolicyUnset struct {
	Name              string `json:"name"`
	Url               bool   `json:"url,string,omitempty"`
	Rule              bool   `json:"rule,string,omitempty"`
	Delay             bool   `json:"delay,string,omitempty"`
	Maxconn           bool   `json:"maxconn,string,omitempty"`
	Action            bool   `json:"action,string,omitempty"`
	Altcontentsvcname bool   `json:"altcontentsvcname,string,omitempty"`
	Altcontentpath    bool   `json:"altcontentpath,string,omitempty"`
}

type update_scpolicy struct {
	Name              string `json:"name"`
	Url               string `json:"url,omitempty"`
	Rule              string `json:"rule,omitempty"`
	Delay             int    `json:"delay,string,omitempty"`
	Maxconn           int    `json:"maxconn,string,omitempty"`
	Action            string `json:"action,omitempty"`
	Altcontentsvcname string `json:"altcontentsvcname,omitempty"`
	Altcontentpath    string `json:"altcontentpath,omitempty"`
}

type rename_scpolicy struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_scpolicy_payload struct {
	Resource Scpolicy `json:"scpolicy"`
}

type rename_scpolicy_payload struct {
	Rename rename_scpolicy `json:"scpolicy"`
}

type unset_scpolicy_payload struct {
	Unset ScpolicyUnset `json:"scpolicy"`
}

type update_scpolicy_payload struct {
	Update update_scpolicy `json:"scpolicy"`
}

type get_scpolicy_result struct {
	Results []Scpolicy `json:"scpolicy"`
}

type count_scpolicy_result struct {
	Results []Count `json:"scpolicy"`
}

func (c *NitroClient) AddScpolicy(resource Scpolicy) error {
	payload := add_scpolicy_payload{
		resource,
	}

	return c.post("scpolicy", "", nil, payload)
}

func (c *NitroClient) RenameScpolicy(name string, newName string) error {
	payload := rename_scpolicy_payload{
		rename_scpolicy{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("scpolicy", "", qs, payload)
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

func (c *NitroClient) ExistsScpolicy(key string) (bool, error) {
	var results count_scpolicy_result

	id, qs := scpolicy_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("scpolicy", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListScpolicy() ([]Scpolicy, error) {
	results := get_scpolicy_result{}

	if err := c.get("scpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetScpolicy(key string) (*Scpolicy, error) {
	var results get_scpolicy_result

	id, qs := scpolicy_key_to_id_args(key)

	if err := c.get("scpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one scpolicy element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("scpolicy element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteScpolicy(key string) error {
	id, qs := scpolicy_key_to_id_args(key)

	return c.delete("scpolicy", id, qs)
}

func (c *NitroClient) UnsetScpolicy(unset ScpolicyUnset) error {
	payload := unset_scpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("scpolicy", "", qs, payload)
}

func (c *NitroClient) UpdateScpolicy(resource Scpolicy) error {
	payload := update_scpolicy_payload{
		update_scpolicy{
			resource.Name,
			resource.Url,
			resource.Rule,
			resource.Delay,
			resource.Maxconn,
			resource.Action,
			resource.Altcontentsvcname,
			resource.Altcontentpath,
		},
	}

	return c.put("scpolicy", "", nil, payload)
}
