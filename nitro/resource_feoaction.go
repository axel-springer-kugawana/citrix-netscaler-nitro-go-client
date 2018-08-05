package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Feoaction struct {
	Name                   string   `json:"name"`
	Cachemaxage            int      `json:"cachemaxage,string,omitempty"`
	Clientsidemeasurements bool     `json:"clientsidemeasurements,omitempty"`
	Convertimporttolink    bool     `json:"convertimporttolink,omitempty"`
	Csscombine             bool     `json:"csscombine,omitempty"`
	Cssimginline           bool     `json:"cssimginline,omitempty"`
	Cssinline              bool     `json:"cssinline,omitempty"`
	Cssminify              bool     `json:"cssminify,omitempty"`
	Cssmovetohead          bool     `json:"cssmovetohead,omitempty"`
	Dnsshards              []string `json:"dnsshards,omitempty"`
	Domainsharding         string   `json:"domainsharding,omitempty"`
	Htmlminify             bool     `json:"htmlminify,omitempty"`
	Imggiftopng            bool     `json:"imggiftopng,omitempty"`
	Imginline              bool     `json:"imginline,omitempty"`
	Imglazyload            bool     `json:"imglazyload,omitempty"`
	Imgshrinktoattrib      bool     `json:"imgshrinktoattrib,omitempty"`
	Imgtojpegxr            bool     `json:"imgtojpegxr,omitempty"`
	Imgtowebp              bool     `json:"imgtowebp,omitempty"`
	Jpgoptimize            bool     `json:"jpgoptimize,omitempty"`
	Jsinline               bool     `json:"jsinline,omitempty"`
	Jsminify               bool     `json:"jsminify,omitempty"`
	Jsmovetoend            bool     `json:"jsmovetoend,omitempty"`
	Pageextendcache        bool     `json:"pageextendcache,omitempty"`
}

func feoaction_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type FeoactionUnset struct {
	Name                   string `json:"name"`
	Pageextendcache        bool   `json:"pageextendcache,omitempty"`
	Cachemaxage            bool   `json:"cachemaxage,omitempty"`
	Imgshrinktoattrib      bool   `json:"imgshrinktoattrib,omitempty"`
	Imggiftopng            bool   `json:"imggiftopng,omitempty"`
	Imgtowebp              bool   `json:"imgtowebp,omitempty"`
	Imgtojpegxr            bool   `json:"imgtojpegxr,omitempty"`
	Imginline              bool   `json:"imginline,omitempty"`
	Cssimginline           bool   `json:"cssimginline,omitempty"`
	Jpgoptimize            bool   `json:"jpgoptimize,omitempty"`
	Imglazyload            bool   `json:"imglazyload,omitempty"`
	Cssminify              bool   `json:"cssminify,omitempty"`
	Cssinline              bool   `json:"cssinline,omitempty"`
	Csscombine             bool   `json:"csscombine,omitempty"`
	Convertimporttolink    bool   `json:"convertimporttolink,omitempty"`
	Jsminify               bool   `json:"jsminify,omitempty"`
	Jsinline               bool   `json:"jsinline,omitempty"`
	Htmlminify             bool   `json:"htmlminify,omitempty"`
	Cssmovetohead          bool   `json:"cssmovetohead,omitempty"`
	Jsmovetoend            bool   `json:"jsmovetoend,omitempty"`
	Domainsharding         bool   `json:"domainsharding,omitempty"`
	Dnsshards              bool   `json:"dnsshards,omitempty"`
	Clientsidemeasurements bool   `json:"clientsidemeasurements,omitempty"`
}

type update_feoaction struct {
	Name                   string   `json:"name"`
	Pageextendcache        bool     `json:"pageextendcache,omitempty"`
	Cachemaxage            int      `json:"cachemaxage,string,omitempty"`
	Imgshrinktoattrib      bool     `json:"imgshrinktoattrib,omitempty"`
	Imggiftopng            bool     `json:"imggiftopng,omitempty"`
	Imgtowebp              bool     `json:"imgtowebp,omitempty"`
	Imgtojpegxr            bool     `json:"imgtojpegxr,omitempty"`
	Imginline              bool     `json:"imginline,omitempty"`
	Cssimginline           bool     `json:"cssimginline,omitempty"`
	Jpgoptimize            bool     `json:"jpgoptimize,omitempty"`
	Imglazyload            bool     `json:"imglazyload,omitempty"`
	Cssminify              bool     `json:"cssminify,omitempty"`
	Cssinline              bool     `json:"cssinline,omitempty"`
	Csscombine             bool     `json:"csscombine,omitempty"`
	Convertimporttolink    bool     `json:"convertimporttolink,omitempty"`
	Jsminify               bool     `json:"jsminify,omitempty"`
	Jsinline               bool     `json:"jsinline,omitempty"`
	Htmlminify             bool     `json:"htmlminify,omitempty"`
	Cssmovetohead          bool     `json:"cssmovetohead,omitempty"`
	Jsmovetoend            bool     `json:"jsmovetoend,omitempty"`
	Domainsharding         string   `json:"domainsharding,omitempty"`
	Dnsshards              []string `json:"dnsshards,omitempty"`
	Clientsidemeasurements bool     `json:"clientsidemeasurements,omitempty"`
}

type rename_feoaction struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_feoaction_payload struct {
	Resource Feoaction `json:"feoaction"`
}

type rename_feoaction_payload struct {
	Rename rename_feoaction `json:"feoaction"`
}

type unset_feoaction_payload struct {
	Unset FeoactionUnset `json:"feoaction"`
}

type update_feoaction_payload struct {
	Update update_feoaction `json:"feoaction"`
}

type get_feoaction_result struct {
	Results []Feoaction `json:"feoaction"`
}

type count_feoaction_result struct {
	Results []Count `json:"feoaction"`
}

func (c *NitroClient) AddFeoaction(resource Feoaction) error {
	payload := add_feoaction_payload{
		resource,
	}

	return c.post("feoaction", "", nil, payload)
}

func (c *NitroClient) RenameFeoaction(name string, newName string) error {
	payload := rename_feoaction_payload{
		rename_feoaction{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("feoaction", "", qs, payload)
}

func (c *NitroClient) CountFeoaction() (int, error) {
	var results count_feoaction_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("feoaction", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsFeoaction(key string) (bool, error) {
	var results count_feoaction_result

	id, qs := feoaction_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("feoaction", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListFeoaction() ([]Feoaction, error) {
	results := get_feoaction_result{}

	if err := c.get("feoaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetFeoaction(key string) (*Feoaction, error) {
	var results get_feoaction_result

	id, qs := feoaction_key_to_id_args(key)

	if err := c.get("feoaction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one feoaction element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("feoaction element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteFeoaction(key string) error {
	id, qs := feoaction_key_to_id_args(key)

	return c.delete("feoaction", id, qs)
}

func (c *NitroClient) UnsetFeoaction(unset FeoactionUnset) error {
	payload := unset_feoaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("feoaction", "", qs, payload)
}

func (c *NitroClient) UpdateFeoaction(resource Feoaction) error {
	payload := update_feoaction_payload{
		update_feoaction{
			resource.Name,
			resource.Pageextendcache,
			resource.Cachemaxage,
			resource.Imgshrinktoattrib,
			resource.Imggiftopng,
			resource.Imgtowebp,
			resource.Imgtojpegxr,
			resource.Imginline,
			resource.Cssimginline,
			resource.Jpgoptimize,
			resource.Imglazyload,
			resource.Cssminify,
			resource.Cssinline,
			resource.Csscombine,
			resource.Convertimporttolink,
			resource.Jsminify,
			resource.Jsinline,
			resource.Htmlminify,
			resource.Cssmovetohead,
			resource.Jsmovetoend,
			resource.Domainsharding,
			resource.Dnsshards,
			resource.Clientsidemeasurements,
		},
	}

	return c.put("feoaction", "", nil, payload)
}
