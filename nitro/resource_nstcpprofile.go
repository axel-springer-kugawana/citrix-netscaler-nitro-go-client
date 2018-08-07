package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Nstcpprofile struct {
	Ackaggregation              string `json:"ackaggregation,omitempty"`
	Ackonpush                   string `json:"ackonpush,omitempty"`
	Buffersize                  int    `json:"buffersize,string,omitempty"`
	Burstratecontrol            string `json:"burstratecontrol,omitempty"`
	Delayedack                  int    `json:"delayedack,string,omitempty"`
	Dropestconnontimeout        string `json:"dropestconnontimeout,omitempty"`
	Drophalfclosedconnontimeout string `json:"drophalfclosedconnontimeout,omitempty"`
	Dsack                       string `json:"dsack,omitempty"`
	Dupackthresh                int    `json:"dupackthresh,string,omitempty"`
	Dynamicreceivebuffering     string `json:"dynamicreceivebuffering,omitempty"`
	Ecn                         string `json:"ecn,omitempty"`
	Establishclientconn         string `json:"establishclientconn,omitempty"`
	Fack                        string `json:"fack,omitempty"`
	Flavor                      string `json:"flavor,omitempty"`
	Frto                        string `json:"frto,omitempty"`
	Hystart                     string `json:"hystart,omitempty"`
	Initialcwnd                 int    `json:"initialcwnd,string,omitempty"`
	Ka                          string `json:"ka,omitempty"`
	Kaconnidletime              int    `json:"kaconnidletime,string,omitempty"`
	Kamaxprobes                 int    `json:"kamaxprobes,string,omitempty"`
	Kaprobeinterval             int    `json:"kaprobeinterval,string,omitempty"`
	Kaprobeupdatelastactivity   string `json:"kaprobeupdatelastactivity,omitempty"`
	Maxburst                    int    `json:"maxburst,string,omitempty"`
	Maxcwnd                     int    `json:"maxcwnd,string,omitempty"`
	Maxpktpermss                int    `json:"maxpktpermss,string,omitempty"`
	Minrto                      int    `json:"minrto,string,omitempty"`
	Mptcp                       string `json:"mptcp,omitempty"`
	Mptcpdropdataonpreestsf     string `json:"mptcpdropdataonpreestsf,omitempty"`
	Mptcpfastopen               string `json:"mptcpfastopen,omitempty"`
	Mptcpsessiontimeout         int    `json:"mptcpsessiontimeout,string,omitempty"`
	Mss                         int    `json:"mss,string,omitempty"`
	Nagle                       string `json:"nagle,omitempty"`
	Name                        string `json:"name,omitempty"`
	Oooqsize                    int    `json:"oooqsize,string,omitempty"`
	Pktperretx                  int    `json:"pktperretx,string,omitempty"`
	Rateqmax                    int    `json:"rateqmax,string,omitempty"`
	Rstmaxack                   string `json:"rstmaxack,omitempty"`
	Rstwindowattenuate          string `json:"rstwindowattenuate,omitempty"`
	Sack                        string `json:"sack,omitempty"`
	Sendbuffsize                int    `json:"sendbuffsize,string,omitempty"`
	Slowstartincr               int    `json:"slowstartincr,string,omitempty"`
	Spoofsyndrop                string `json:"spoofsyndrop,omitempty"`
	Syncookie                   string `json:"syncookie,omitempty"`
	Tcpfastopen                 string `json:"tcpfastopen,omitempty"`
	Tcpmode                     string `json:"tcpmode,omitempty"`
	Tcprate                     int    `json:"tcprate,string,omitempty"`
	Tcpsegoffload               string `json:"tcpsegoffload,omitempty"`
	Timestamp                   string `json:"timestamp,omitempty"`
	Ws                          string `json:"ws,omitempty"`
	Wsval                       int    `json:"wsval,string,omitempty"`
}

type NstcpprofileKey struct {
	Name string
}

func (resource Nstcpprofile) ToKey() NstcpprofileKey {
	key := NstcpprofileKey{
		resource.Name,
	}

	return key
}

func (key NstcpprofileKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Name

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key NstcpprofileKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key NstcpprofileKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_nstcpprofile_payload struct {
	Resource Nstcpprofile `json:"nstcpprofile"`
}

func (c *NitroClient) AddNstcpprofile(resource Nstcpprofile) error {
	payload := add_nstcpprofile_payload{
		resource,
	}

	return c.post("nstcpprofile", "", nil, payload)
}

//      LIST

type list_nstcpprofile_result struct {
	Results []Nstcpprofile `json:"nstcpprofile"`
}

func (c *NitroClient) ListNstcpprofile() ([]Nstcpprofile, error) {
	results := list_nstcpprofile_result{}

	if err := c.get("nstcpprofile", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_nstcpprofile_result struct {
	Results []Nstcpprofile `json:"nstcpprofile"`
}

func (c *NitroClient) GetNstcpprofile(key NstcpprofileKey) (*Nstcpprofile, error) {
	var results get_nstcpprofile_result

	id, qs := key.to_id_args()

	if err := c.get("nstcpprofile", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one nstcpprofile element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("nstcpprofile element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_nstcpprofile_result struct {
	Results []Count `json:"nstcpprofile"`
}

func (c *NitroClient) CountNstcpprofile() (int, error) {
	var results count_nstcpprofile_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("nstcpprofile", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsNstcpprofile(key NstcpprofileKey) (bool, error) {
	var results count_nstcpprofile_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("nstcpprofile", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteNstcpprofile(key NstcpprofileKey) error {
	id, qs := key.to_id_args()

	return c.delete("nstcpprofile", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
