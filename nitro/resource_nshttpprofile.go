package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Nshttpprofile struct {
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

func nshttpprofile_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type NshttpprofileUnset struct {
	Name                      string `json:"name"`
	Dropinvalreqs             bool   `json:"dropinvalreqs,omitempty"`
	Markhttp09inval           bool   `json:"markhttp09inval,omitempty"`
	Markconnreqinval          bool   `json:"markconnreqinval,omitempty"`
	Cmponpush                 bool   `json:"cmponpush,omitempty"`
	Conmultiplex              bool   `json:"conmultiplex,omitempty"`
	Maxreusepool              bool   `json:"maxreusepool,omitempty"`
	Dropextracrlf             bool   `json:"dropextracrlf,omitempty"`
	Incomphdrdelay            bool   `json:"incomphdrdelay,omitempty"`
	Websocket                 bool   `json:"websocket,omitempty"`
	Rtsptunnel                bool   `json:"rtsptunnel,omitempty"`
	Reqtimeout                bool   `json:"reqtimeout,omitempty"`
	Adpttimeout               bool   `json:"adpttimeout,omitempty"`
	Reqtimeoutaction          bool   `json:"reqtimeoutaction,omitempty"`
	Dropextradata             bool   `json:"dropextradata,omitempty"`
	Weblog                    bool   `json:"weblog,omitempty"`
	Clientiphdrexpr           bool   `json:"clientiphdrexpr,omitempty"`
	Maxreq                    bool   `json:"maxreq,omitempty"`
	Persistentetag            bool   `json:"persistentetag,omitempty"`
	Spdy                      bool   `json:"spdy,omitempty"`
	Http2                     bool   `json:"http2,omitempty"`
	Http2direct               bool   `json:"http2direct,omitempty"`
	Altsvc                    bool   `json:"altsvc,omitempty"`
	Http2maxheaderlistsize    bool   `json:"http2maxheaderlistsize,omitempty"`
	Http2maxframesize         bool   `json:"http2maxframesize,omitempty"`
	Http2maxconcurrentstreams bool   `json:"http2maxconcurrentstreams,omitempty"`
	Http2initialwindowsize    bool   `json:"http2initialwindowsize,omitempty"`
	Http2headertablesize      bool   `json:"http2headertablesize,omitempty"`
	Http2minseverconn         bool   `json:"http2minseverconn,omitempty"`
	Reusepooltimeout          bool   `json:"reusepooltimeout,omitempty"`
	Maxheaderlen              bool   `json:"maxheaderlen,omitempty"`
	Minreusepool              bool   `json:"minreusepool,omitempty"`
	Apdexcltresptimethreshold bool   `json:"apdexcltresptimethreshold,omitempty"`
}

type update_nshttpprofile struct {
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

type rename_nshttpprofile struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_nshttpprofile_payload struct {
	Resource Nshttpprofile `json:"nshttpprofile"`
}

type rename_nshttpprofile_payload struct {
	Rename rename_nshttpprofile `json:"nshttpprofile"`
}

type unset_nshttpprofile_payload struct {
	Unset NshttpprofileUnset `json:"nshttpprofile"`
}

type update_nshttpprofile_payload struct {
	Update update_nshttpprofile `json:"nshttpprofile"`
}

type get_nshttpprofile_result struct {
	Results []Nshttpprofile `json:"nshttpprofile"`
}

type count_nshttpprofile_result struct {
	Results []Count `json:"nshttpprofile"`
}

func (c *NitroClient) AddNshttpprofile(resource Nshttpprofile) error {
	payload := add_nshttpprofile_payload{
		resource,
	}

	return c.post("nshttpprofile", "", nil, payload)
}

func (c *NitroClient) RenameNshttpprofile(name string, newName string) error {
	payload := rename_nshttpprofile_payload{
		rename_nshttpprofile{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("nshttpprofile", "", qs, payload)
}

func (c *NitroClient) CountNshttpprofile() (int, error) {
	var results count_nshttpprofile_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("nshttpprofile", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsNshttpprofile(key string) (bool, error) {
	var results count_nshttpprofile_result

	id, qs := nshttpprofile_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("nshttpprofile", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListNshttpprofile() ([]Nshttpprofile, error) {
	results := get_nshttpprofile_result{}

	if err := c.get("nshttpprofile", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetNshttpprofile(key string) (*Nshttpprofile, error) {
	var results get_nshttpprofile_result

	id, qs := nshttpprofile_key_to_id_args(key)

	if err := c.get("nshttpprofile", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one nshttpprofile element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("nshttpprofile element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteNshttpprofile(key string) error {
	id, qs := nshttpprofile_key_to_id_args(key)

	return c.delete("nshttpprofile", id, qs)
}

func (c *NitroClient) UnsetNshttpprofile(unset NshttpprofileUnset) error {
	payload := unset_nshttpprofile_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("nshttpprofile", "", qs, payload)
}

func (c *NitroClient) UpdateNshttpprofile(resource Nshttpprofile) error {
	payload := update_nshttpprofile_payload{
		update_nshttpprofile{
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
		},
	}

	return c.put("nshttpprofile", "", nil, payload)
}
