package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Cachepolicylabel struct {
	Labelname string `json:"labelname"`
	Evaluates string `json:"evaluates,omitempty"`
}

func cachepolicylabel_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	qs := map[string]string{}

	return key, qs
}

type rename_cachepolicylabel struct {
	Name    string `json:"labelname"`
	Newname string `json:"newname"`
}

type add_cachepolicylabel_payload struct {
	Resource Cachepolicylabel `json:"cachepolicylabel"`
}

type rename_cachepolicylabel_payload struct {
	Rename rename_cachepolicylabel `json:"cachepolicylabel"`
}

type get_cachepolicylabel_result struct {
	Results []Cachepolicylabel `json:"cachepolicylabel"`
}

type count_cachepolicylabel_result struct {
	Results []Count `json:"cachepolicylabel"`
}

func (c *NitroClient) AddCachepolicylabel(resource Cachepolicylabel) error {
	payload := add_cachepolicylabel_payload{
		resource,
	}

	return c.post("cachepolicylabel", "", nil, payload)
}

func (c *NitroClient) RenameCachepolicylabel(name string, newName string) error {
	payload := rename_cachepolicylabel_payload{
		rename_cachepolicylabel{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("cachepolicylabel", "", qs, payload)
}

func (c *NitroClient) CountCachepolicylabel() (int, error) {
	var results count_cachepolicylabel_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("cachepolicylabel", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsCachepolicylabel(key string) (bool, error) {
	var results count_cachepolicylabel_result

	id, qs := cachepolicylabel_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("cachepolicylabel", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListCachepolicylabel() ([]Cachepolicylabel, error) {
	results := get_cachepolicylabel_result{}

	if err := c.get("cachepolicylabel", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetCachepolicylabel(key string) (*Cachepolicylabel, error) {
	var results get_cachepolicylabel_result

	id, qs := cachepolicylabel_key_to_id_args(key)

	if err := c.get("cachepolicylabel", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one cachepolicylabel element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("cachepolicylabel element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteCachepolicylabel(key string) error {
	id, qs := cachepolicylabel_key_to_id_args(key)

	return c.delete("cachepolicylabel", id, qs)
}
