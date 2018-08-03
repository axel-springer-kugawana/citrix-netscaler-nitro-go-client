package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Appflowpolicylabel struct {
	Labelname       string `json:"labelname"`
	Policylabeltype string `json:"policylabeltype,omitempty"`
}

func appflowpolicylabel_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type rename_appflowpolicylabel struct {
	Name    string `json:"labelname"`
	Newname string `json:"newname"`
}

type add_appflowpolicylabel_payload struct {
	Resource Appflowpolicylabel `json:"appflowpolicylabel"`
}

type rename_appflowpolicylabel_payload struct {
	Rename rename_appflowpolicylabel `json:"appflowpolicylabel"`
}

type get_appflowpolicylabel_result struct {
	Results []Appflowpolicylabel `json:"appflowpolicylabel"`
}

type count_appflowpolicylabel_result struct {
	Results []Count `json:"appflowpolicylabel"`
}

func (c *NitroClient) AddAppflowpolicylabel(resource Appflowpolicylabel) error {
	payload := add_appflowpolicylabel_payload{
		resource,
	}

	return c.post("appflowpolicylabel", "", nil, payload)
}

func (c *NitroClient) RenameAppflowpolicylabel(name string, newName string) error {
	payload := rename_appflowpolicylabel_payload{
		rename_appflowpolicylabel{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("appflowpolicylabel", "", qs, payload)
}

func (c *NitroClient) CountAppflowpolicylabel() (int, error) {
	var results count_appflowpolicylabel_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("appflowpolicylabel", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsAppflowpolicylabel(key string) (bool, error) {
	var results count_appflowpolicylabel_result

	id, qs := appflowpolicylabel_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("appflowpolicylabel", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListAppflowpolicylabel() ([]Appflowpolicylabel, error) {
	results := get_appflowpolicylabel_result{}

	if err := c.get("appflowpolicylabel", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetAppflowpolicylabel(key string) (*Appflowpolicylabel, error) {
	var results get_appflowpolicylabel_result

	id, qs := appflowpolicylabel_key_to_id_args(key)

	if err := c.get("appflowpolicylabel", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one appflowpolicylabel element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("appflowpolicylabel element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteAppflowpolicylabel(key string) error {
	id, qs := appflowpolicylabel_key_to_id_args(key)

	return c.delete("appflowpolicylabel", id, qs)
}
