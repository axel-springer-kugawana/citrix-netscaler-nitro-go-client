package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Feoaction struct {
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
	Name                   string   `json:"name,omitempty"`
	Pageextendcache        bool     `json:"pageextendcache,omitempty"`
}

type FeoactionKey struct {
	Name string
}

func (resource Feoaction) ToKey() FeoactionKey {
	key := FeoactionKey{
		resource.Name,
	}

	return key
}

func (key FeoactionKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = key.Name

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key FeoactionKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key FeoactionKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_feoaction_payload struct {
	Resource Feoaction `json:"feoaction"`
}

func (c *NitroClient) AddFeoaction(resource Feoaction) error {
	payload := add_feoaction_payload{
		resource,
	}

	return c.post("feoaction", "", nil, payload)
}

//      LIST

type list_feoaction_result struct {
	Results []Feoaction `json:"feoaction"`
}

func (c *NitroClient) ListFeoaction() ([]Feoaction, error) {
	results := list_feoaction_result{}

	if err := c.get("feoaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_feoaction_result struct {
	Results []Feoaction `json:"feoaction"`
}

func (c *NitroClient) GetFeoaction(key FeoactionKey) (*Feoaction, error) {
	var results get_feoaction_result

	id, qs := key.to_id_args()

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

//      COUNT

type count_feoaction_result struct {
	Results []Count `json:"feoaction"`
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

//      EXISTS

func (c *NitroClient) ExistsFeoaction(key FeoactionKey) (bool, error) {
	var results count_feoaction_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("feoaction", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteFeoaction(key FeoactionKey) error {
	id, qs := key.to_id_args()

	return c.delete("feoaction", id, qs)
}

//      UPDATE

type FeoactionUpdate struct {
	Name                   string   `json:"name,omitempty"`
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

func (resource Feoaction) ToUpdate() FeoactionUpdate {
	update := FeoactionUpdate{
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
	}

	return update
}

type update_feoaction_payload struct {
	Update FeoactionUpdate `json:"feoaction"`
}

func (c *NitroClient) UpdateFeoaction(update FeoactionUpdate) error {
	payload := update_feoaction_payload{
		update,
	}

	return c.put("feoaction", "", nil, payload)
}

//      UNSET

type FeoactionUnset struct {
	Name                   string `json:"name,omitempty"`
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

func (resource Feoaction) ToUnset() FeoactionUnset {
	unset := FeoactionUnset{
		resource.Name,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
	}

	return unset
}

type unset_feoaction_payload struct {
	Unset FeoactionUnset `json:"feoaction"`
}

func (c *NitroClient) UnsetFeoaction(unset FeoactionUnset) error {
	payload := unset_feoaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("feoaction", "", qs, payload)
}

//      RENAME
//      TODO
