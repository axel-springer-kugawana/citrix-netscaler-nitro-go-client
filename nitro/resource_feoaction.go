package nitro

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

type FeoactionKey struct {
	Name string
}

type feoaction_update struct {
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

type feoaction_payload struct {
	feoaction interface{}
}

func feoaction_key_to_args(key FeoactionKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteFeoaction(key FeoactionKey) error {
	return c.deleteResourceWithArgs("feoaction", key.Name, feoaction_key_to_args(key))
}

func (c *NitroClient) GetFeoaction(key FeoactionKey) (*Feoaction, error) {
	var results struct {
		Feoaction []Feoaction
	}

	if err := c.getResourceWithArgs("feoaction", key.Name, feoaction_key_to_args(key), &results); err != nil || len(results.Feoaction) != 1 {
		return nil, err
	}

	return &results.Feoaction[0], nil
}

func (c *NitroClient) ListFeoaction() ([]Feoaction, error) {
	var results struct {
		Feoaction []Feoaction
	}

	if err := c.listResources("feoaction", &results); err != nil {
		return nil, err
	}

	return results.Feoaction, nil
}

func (c *NitroClient) AddFeoaction(resource Feoaction) error {
	return c.addResource("feoaction", resource)
}

func (c *NitroClient) RenameFeoaction(name string, newName string) error {
	return c.renameResource("feoaction", "name", name, newName)
}

func (c *NitroClient) UnsetFeoaction(name string, fields ...string) error {
	return c.unsetResource("feoaction", "name", name, fields)
}

func (c *NitroClient) UpdateFeoaction(resource Feoaction) error {
	update := feoaction_update{
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

	return c.Put("feoaction", update)
}
