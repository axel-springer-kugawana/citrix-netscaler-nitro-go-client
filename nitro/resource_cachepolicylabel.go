package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Cachepolicylabel struct {
	Evaluates string `json:"evaluates,omitempty"`
	Labelname string `json:"labelname,omitempty"`
}

type CachepolicylabelKey struct {
	Labelname string
}

func (resource Cachepolicylabel) ToKey() CachepolicylabelKey {
	key := CachepolicylabelKey{
		resource.Labelname,
	}

	return key
}

func (key CachepolicylabelKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = key.Labelname

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key CachepolicylabelKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key CachepolicylabelKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_cachepolicylabel_payload struct {
	Resource Cachepolicylabel `json:"cachepolicylabel"`
}

func (c *NitroClient) AddCachepolicylabel(resource Cachepolicylabel) error {
	payload := add_cachepolicylabel_payload{
		resource,
	}

	return c.post("cachepolicylabel", "", nil, payload)
}

//      LIST

type list_cachepolicylabel_result struct {
	Results []Cachepolicylabel `json:"cachepolicylabel"`
}

func (c *NitroClient) ListCachepolicylabel() ([]Cachepolicylabel, error) {
	results := list_cachepolicylabel_result{}

	if err := c.get("cachepolicylabel", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_cachepolicylabel_result struct {
	Results []Cachepolicylabel `json:"cachepolicylabel"`
}

func (c *NitroClient) GetCachepolicylabel(key CachepolicylabelKey) (*Cachepolicylabel, error) {
	var results get_cachepolicylabel_result

	id, qs := key.to_id_args()

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

//      COUNT

type count_cachepolicylabel_result struct {
	Results []Count `json:"cachepolicylabel"`
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

//      EXISTS

func (c *NitroClient) ExistsCachepolicylabel(key CachepolicylabelKey) (bool, error) {
	var results count_cachepolicylabel_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("cachepolicylabel", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteCachepolicylabel(key CachepolicylabelKey) error {
	id, qs := key.to_id_args()

	return c.delete("cachepolicylabel", id, qs)
}

//      RENAME
//      TODO
