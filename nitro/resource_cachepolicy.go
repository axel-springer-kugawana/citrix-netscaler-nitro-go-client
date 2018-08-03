package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Cachepolicy struct {
	Policyname   string   `json:"policyname"`
	Action       string   `json:"action,omitempty"`
	Invalgroups  []string `json:"invalgroups,omitempty"`
	Invalobjects []string `json:"invalobjects,omitempty"`
	Rule         string   `json:"rule,omitempty"`
	Storeingroup string   `json:"storeingroup,omitempty"`
	Undefaction  string   `json:"undefaction,omitempty"`
}

type CachepolicyKey struct {
	Policyname string `json:"policyname"`
}

type CachepolicyUnset struct {
	Policyname   string `json:"policyname"`
	Rule         bool   `json:"rule,string,omitempty"`
	Action       bool   `json:"action,string,omitempty"`
	Storeingroup bool   `json:"storeingroup,string,omitempty"`
	Invalgroups  bool   `json:"invalgroups,string,omitempty"`
	Invalobjects bool   `json:"invalobjects,string,omitempty"`
	Undefaction  bool   `json:"undefaction,string,omitempty"`
}

type update_cachepolicy struct {
	Policyname   string   `json:"policyname"`
	Rule         string   `json:"rule,omitempty"`
	Action       string   `json:"action,omitempty"`
	Storeingroup string   `json:"storeingroup,omitempty"`
	Invalgroups  []string `json:"invalgroups,omitempty"`
	Invalobjects []string `json:"invalobjects,omitempty"`
	Undefaction  string   `json:"undefaction,omitempty"`
}

type rename_cachepolicy struct {
	Name    string `json:"policyname"`
	Newname string `json:"newname"`
}

type add_cachepolicy_payload struct {
	Resource Cachepolicy `json:"cachepolicy"`
}

type rename_cachepolicy_payload struct {
	Rename rename_cachepolicy `json:"cachepolicy"`
}

type unset_cachepolicy_payload struct {
	Unset CachepolicyUnset `json:"cachepolicy"`
}

type update_cachepolicy_payload struct {
	Update update_cachepolicy `json:"cachepolicy"`
}

type get_cachepolicy_result struct {
	Results []Cachepolicy `json:"cachepolicy"`
}

type count_cachepolicy_result struct {
	Results []Count `json:"cachepolicy"`
}

func cachepolicy_key_to_id_args(key CachepolicyKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return key.Policyname, qs
}

func (c *NitroClient) AddCachepolicy(resource Cachepolicy) error {
	payload := add_cachepolicy_payload{
		resource,
	}

	return c.post("cachepolicy", "", nil, payload)
}

func (c *NitroClient) RenameCachepolicy(name string, newName string) error {
	payload := rename_cachepolicy_payload{
		rename_cachepolicy{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("cachepolicy", "", qs, payload)
}

func (c *NitroClient) CountCachepolicy() (int, error) {
	var results count_cachepolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("cachepolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsCachepolicy(key CachepolicyKey) (bool, error) {
	var results count_cachepolicy_result

	id, qs := cachepolicy_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("cachepolicy", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListCachepolicy() ([]Cachepolicy, error) {
	var results get_cachepolicy_result

	if err := c.get("cachepolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetCachepolicy(key CachepolicyKey) (*Cachepolicy, error) {
	var results get_cachepolicy_result

	id, qs := cachepolicy_key_to_id_args(key)

	if err := c.get("cachepolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one cachepolicy element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("cachepolicy element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteCachepolicy(key CachepolicyKey) error {
	id, qs := cachepolicy_key_to_id_args(key)

	return c.delete("cachepolicy", id, qs)
}

func (c *NitroClient) UnsetCachepolicy(unset CachepolicyUnset) error {
	payload := unset_cachepolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("cachepolicy", "", qs, payload)
}

func (c *NitroClient) UpdateCachepolicy(resource Cachepolicy) error {
	payload := update_cachepolicy_payload{
		update_cachepolicy{
			resource.Policyname,
			resource.Rule,
			resource.Action,
			resource.Storeingroup,
			resource.Invalgroups,
			resource.Invalobjects,
			resource.Undefaction,
		},
	}

	return c.put("cachepolicy", "", nil, payload)
}
