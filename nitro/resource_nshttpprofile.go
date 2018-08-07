package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Nshttpprofile struct {
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
	Name                      string `json:"name,omitempty"`
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

func (resource Nshttpprofile) ToKey() NshttpprofileKey {
	key := NshttpprofileKey{
		resource.Name,
	}

	return key
}

func (key NshttpprofileKey) to_id_args() (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Name

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return id, qs
}

//      CREATE

type add_nshttpprofile_payload struct {
	Resource Nshttpprofile `json:"nshttpprofile"`
}

func (c *NitroClient) AddNshttpprofile(resource Nshttpprofile) error {
	payload := add_nshttpprofile_payload{
		resource,
	}

	return c.post("nshttpprofile", "", nil, payload)
}

//      LIST

type list_nshttpprofile_result struct {
	Results []Nshttpprofile `json:"nshttpprofile"`
}

func (c *NitroClient) ListNshttpprofile() ([]Nshttpprofile, error) {
	results := list_nshttpprofile_result{}

	if err := c.get("nshttpprofile", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_nshttpprofile_result struct {
	Results []Nshttpprofile `json:"nshttpprofile"`
}

func (c *NitroClient) GetNshttpprofile(key NshttpprofileKey) (*Nshttpprofile, error) {
	var results get_nshttpprofile_result

	id, qs := key.to_id_args()

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

//      COUNT

type count_nshttpprofile_result struct {
	Results []Count `json:"nshttpprofile"`
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

//      EXISTS

func (c *NitroClient) ExistsNshttpprofile(key NshttpprofileKey) (bool, error) {
	var results count_nshttpprofile_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("nshttpprofile", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteNshttpprofile(key NshttpprofileKey) error {
	id, qs := key.to_id_args()

	return c.delete("nshttpprofile", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
