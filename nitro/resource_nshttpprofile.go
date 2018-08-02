package nitro

type NshttpprofileResource struct {
	Name                      string `json:"name"`
	Adpttimeout               string `json:"adpttimeout,omitempty"`
	Altsvc                    string `json:"altsvc,omitempty"`
	Apdexcltresptimethreshold int    `json:"apdexcltresptimethreshold,string,omitempty"`
	Clientiphdrexpr           string `json:"clientiphdrexpr,omitempty"`
	Cmponpush                 string `json:"cmponpush,omitempty"`
	Conmultiplex              string `json:"conmultiplex,omitempty"`
	Dropextracrlf             string `json:"dropextracrlf,omitempty"`
	Dropextradata             string `json:"dropextradata,omitempty"`
	Dropinvalreqs             string `json:"dropinvalreqs,omitempty"`
	Http2                     string `json:"http2,omitempty"`
	Http2direct               string `json:"http2direct,omitempty"`
	Http2headertablesize      int    `json:"http2headertablesize,string,omitempty"`
	Http2initialwindowsize    int    `json:"http2initialwindowsize,string,omitempty"`
	Http2maxconcurrentstreams int    `json:"http2maxconcurrentstreams,string,omitempty"`
	Http2maxframesize         int    `json:"http2maxframesize,string,omitempty"`
	Http2maxheaderlistsize    int    `json:"http2maxheaderlistsize,string,omitempty"`
	Http2minseverconn         int    `json:"http2minseverconn,string,omitempty"`
	Incomphdrdelay            int    `json:"incomphdrdelay,string,omitempty"`
	Markconnreqinval          string `json:"markconnreqinval,omitempty"`
	Markhttp09inval           string `json:"markhttp09inval,omitempty"`
	Maxheaderlen              int    `json:"maxheaderlen,string,omitempty"`
	Maxreq                    int    `json:"maxreq,string,omitempty"`
	Maxreusepool              int    `json:"maxreusepool,string,omitempty"`
	Minreusepool              int    `json:"minreusepool,string,omitempty"`
	Persistentetag            string `json:"persistentetag,omitempty"`
	Reqtimeout                int    `json:"reqtimeout,string,omitempty"`
	Reqtimeoutaction          string `json:"reqtimeoutaction,omitempty"`
	Reusepooltimeout          int    `json:"reusepooltimeout,string,omitempty"`
	Rtsptunnel                string `json:"rtsptunnel,omitempty"`
	Spdy                      string `json:"spdy,omitempty"`
	Weblog                    string `json:"weblog,omitempty"`
	Websocket                 string `json:"websocket,omitempty"`
}

type NshttpprofileKey struct {
	Name string
}

type nshttpprofile_update struct {
	Name                      string `json:"name"`
	Dropinvalreqs             string `json:"dropinvalreqs,omitempty"`
	Markhttp09inval           string `json:"markhttp09inval,omitempty"`
	Markconnreqinval          string `json:"markconnreqinval,omitempty"`
	Cmponpush                 string `json:"cmponpush,omitempty"`
	Conmultiplex              string `json:"conmultiplex,omitempty"`
	Maxreusepool              int    `json:"maxreusepool,string,omitempty"`
	Dropextracrlf             string `json:"dropextracrlf,omitempty"`
	Incomphdrdelay            int    `json:"incomphdrdelay,string,omitempty"`
	Websocket                 string `json:"websocket,omitempty"`
	Rtsptunnel                string `json:"rtsptunnel,omitempty"`
	Reqtimeout                int    `json:"reqtimeout,string,omitempty"`
	Adpttimeout               string `json:"adpttimeout,omitempty"`
	Reqtimeoutaction          string `json:"reqtimeoutaction,omitempty"`
	Dropextradata             string `json:"dropextradata,omitempty"`
	Weblog                    string `json:"weblog,omitempty"`
	Clientiphdrexpr           string `json:"clientiphdrexpr,omitempty"`
	Maxreq                    int    `json:"maxreq,string,omitempty"`
	Persistentetag            string `json:"persistentetag,omitempty"`
	Spdy                      string `json:"spdy,omitempty"`
	Http2                     string `json:"http2,omitempty"`
	Http2direct               string `json:"http2direct,omitempty"`
	Altsvc                    string `json:"altsvc,omitempty"`
	Http2maxheaderlistsize    int    `json:"http2maxheaderlistsize,string,omitempty"`
	Http2maxframesize         int    `json:"http2maxframesize,string,omitempty"`
	Http2maxconcurrentstreams int    `json:"http2maxconcurrentstreams,string,omitempty"`
	Http2initialwindowsize    int    `json:"http2initialwindowsize,string,omitempty"`
	Http2headertablesize      int    `json:"http2headertablesize,string,omitempty"`
	Http2minseverconn         int    `json:"http2minseverconn,string,omitempty"`
	Reusepooltimeout          int    `json:"reusepooltimeout,string,omitempty"`
	Maxheaderlen              int    `json:"maxheaderlen,string,omitempty"`
	Minreusepool              int    `json:"minreusepool,string,omitempty"`
	Apdexcltresptimethreshold int    `json:"apdexcltresptimethreshold,string,omitempty"`
}

type nshttpprofile_payload struct {
	nshttpprofile interface{}
}

func nshttpprofile_key_to_args(key NshttpprofileKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteNshttpprofile(key NshttpprofileKey) error {
	return c.deleteResourceWithArgs("nshttpprofile", key.Name, nshttpprofile_key_to_args(key))
}

func (c *NitroClient) GetNshttpprofile(key NshttpprofileKey) (*NshttpprofileResource, error) {
	var results struct {
		Nshttpprofile []NshttpprofileResource
	}

	if err := c.getResourceWithArgs("nshttpprofile", key.Name, nshttpprofile_key_to_args(key), &results); err != nil || len(results.Nshttpprofile) != 1 {
		return nil, err
	}

	return &results.Nshttpprofile[0], nil
}

func (c *NitroClient) ListNshttpprofile() ([]NshttpprofileResource, error) {
	var results struct {
		Nshttpprofile []NshttpprofileResource
	}

	if err := c.listResources("nshttpprofile", &results); err != nil {
		return nil, err
	}

	return results.Nshttpprofile, nil
}

func (c *NitroClient) AddNshttpprofile(resource NshttpprofileResource) error {
	return c.addResource("nshttpprofile", resource)
}

func (c *NitroClient) RenameNshttpprofile(name string, newName string) error {
	return c.renameResource("nshttpprofile", "name", name, newName)
}

func (c *NitroClient) UnsetNshttpprofile(name string, fields ...string) error {
	return c.unsetResource("nshttpprofile", "name", name, fields)
}

func (c *NitroClient) UpdateNshttpprofile(resource NshttpprofileResource) error {
	update := nshttpprofile_update{
		resource.Name,
		resource.Dropinvalreqs,
		resource.Markhttp09inval,
		resource.Markconnreqinval,
		resource.Cmponpush,
		resource.Conmultiplex,
		resource.Maxreusepool,
		resource.Dropextracrlf,
		resource.Incomphdrdelay,
		resource.Websocket,
		resource.Rtsptunnel,
		resource.Reqtimeout,
		resource.Adpttimeout,
		resource.Reqtimeoutaction,
		resource.Dropextradata,
		resource.Weblog,
		resource.Clientiphdrexpr,
		resource.Maxreq,
		resource.Persistentetag,
		resource.Spdy,
		resource.Http2,
		resource.Http2direct,
		resource.Altsvc,
		resource.Http2maxheaderlistsize,
		resource.Http2maxframesize,
		resource.Http2maxconcurrentstreams,
		resource.Http2initialwindowsize,
		resource.Http2headertablesize,
		resource.Http2minseverconn,
		resource.Reusepooltimeout,
		resource.Maxheaderlen,
		resource.Minreusepool,
		resource.Apdexcltresptimethreshold,
	}

	return c.Put("nshttpprofile", update)
}
