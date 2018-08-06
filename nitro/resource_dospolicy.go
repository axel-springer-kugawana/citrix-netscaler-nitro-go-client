package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Dospolicy struct {
	Name          string `json:"name"`
	Cltdetectrate int    `json:"cltdetectrate,string,omitempty"`
	Qdepth        int    `json:"qdepth,string,omitempty"`
}

func dospolicy_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	qs := map[string]string{}

	return key, qs
}

type DospolicyUnset struct {
	Name          string `json:"name"`
	Qdepth        bool   `json:"qdepth,omitempty"`
	Cltdetectrate bool   `json:"cltdetectrate,omitempty"`
}

type update_dospolicy struct {
	Name          string `json:"name"`
	Qdepth        int    `json:"qdepth,string,omitempty"`
	Cltdetectrate int    `json:"cltdetectrate,string,omitempty"`
}

type rename_dospolicy struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_dospolicy_payload struct {
	Resource Dospolicy `json:"dospolicy"`
}

type rename_dospolicy_payload struct {
	Rename rename_dospolicy `json:"dospolicy"`
}

type unset_dospolicy_payload struct {
	Unset DospolicyUnset `json:"dospolicy"`
}

type update_dospolicy_payload struct {
	Update update_dospolicy `json:"dospolicy"`
}

type get_dospolicy_result struct {
	Results []Dospolicy `json:"dospolicy"`
}

type count_dospolicy_result struct {
	Results []Count `json:"dospolicy"`
}

func (c *NitroClient) AddDospolicy(resource Dospolicy) error {
	payload := add_dospolicy_payload{
		resource,
	}

	return c.post("dospolicy", "", nil, payload)
}

func (c *NitroClient) RenameDospolicy(name string, newName string) error {
	payload := rename_dospolicy_payload{
		rename_dospolicy{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("dospolicy", "", qs, payload)
}

func (c *NitroClient) CountDospolicy() (int, error) {
	var results count_dospolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("dospolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsDospolicy(key string) (bool, error) {
	var results count_dospolicy_result

	id, qs := dospolicy_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("dospolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListDospolicy() ([]Dospolicy, error) {
	results := get_dospolicy_result{}

	if err := c.get("dospolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetDospolicy(key string) (*Dospolicy, error) {
	var results get_dospolicy_result

	id, qs := dospolicy_key_to_id_args(key)

	if err := c.get("dospolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one dospolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("dospolicy element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteDospolicy(key string) error {
	id, qs := dospolicy_key_to_id_args(key)

	return c.delete("dospolicy", id, qs)
}

func (c *NitroClient) UnsetDospolicy(unset DospolicyUnset) error {
	payload := unset_dospolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("dospolicy", "", qs, payload)
}

func (c *NitroClient) UpdateDospolicy(resource Dospolicy) error {
	payload := update_dospolicy_payload{
		update_dospolicy{
			resource.Name,
			resource.Qdepth,
			resource.Cltdetectrate,
		},
	}

	return c.put("dospolicy", "", nil, payload)
}
