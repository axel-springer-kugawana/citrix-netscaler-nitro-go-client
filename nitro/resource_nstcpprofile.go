package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Nstcpprofile struct {
	Name                        string `json:"name"`
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

func nstcpprofile_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type NstcpprofileUnset struct {
	Name                        string `json:"name"`
	Ws                          bool   `json:"ws,omitempty"`
	Sack                        bool   `json:"sack,omitempty"`
	Wsval                       bool   `json:"wsval,omitempty"`
	Nagle                       bool   `json:"nagle,omitempty"`
	Ackonpush                   bool   `json:"ackonpush,omitempty"`
	Mss                         bool   `json:"mss,omitempty"`
	Maxburst                    bool   `json:"maxburst,omitempty"`
	Initialcwnd                 bool   `json:"initialcwnd,omitempty"`
	Delayedack                  bool   `json:"delayedack,omitempty"`
	Oooqsize                    bool   `json:"oooqsize,omitempty"`
	Maxpktpermss                bool   `json:"maxpktpermss,omitempty"`
	Pktperretx                  bool   `json:"pktperretx,omitempty"`
	Minrto                      bool   `json:"minrto,omitempty"`
	Slowstartincr               bool   `json:"slowstartincr,omitempty"`
	Buffersize                  bool   `json:"buffersize,omitempty"`
	Syncookie                   bool   `json:"syncookie,omitempty"`
	Kaprobeupdatelastactivity   bool   `json:"kaprobeupdatelastactivity,omitempty"`
	Flavor                      bool   `json:"flavor,omitempty"`
	Dynamicreceivebuffering     bool   `json:"dynamicreceivebuffering,omitempty"`
	Ka                          bool   `json:"ka,omitempty"`
	Kaconnidletime              bool   `json:"kaconnidletime,omitempty"`
	Kamaxprobes                 bool   `json:"kamaxprobes,omitempty"`
	Kaprobeinterval             bool   `json:"kaprobeinterval,omitempty"`
	Sendbuffsize                bool   `json:"sendbuffsize,omitempty"`
	Mptcp                       bool   `json:"mptcp,omitempty"`
	Establishclientconn         bool   `json:"establishclientconn,omitempty"`
	Tcpsegoffload               bool   `json:"tcpsegoffload,omitempty"`
	Rstwindowattenuate          bool   `json:"rstwindowattenuate,omitempty"`
	Rstmaxack                   bool   `json:"rstmaxack,omitempty"`
	Spoofsyndrop                bool   `json:"spoofsyndrop,omitempty"`
	Ecn                         bool   `json:"ecn,omitempty"`
	Mptcpdropdataonpreestsf     bool   `json:"mptcpdropdataonpreestsf,omitempty"`
	Mptcpfastopen               bool   `json:"mptcpfastopen,omitempty"`
	Mptcpsessiontimeout         bool   `json:"mptcpsessiontimeout,omitempty"`
	Timestamp                   bool   `json:"timestamp,omitempty"`
	Dsack                       bool   `json:"dsack,omitempty"`
	Ackaggregation              bool   `json:"ackaggregation,omitempty"`
	Frto                        bool   `json:"frto,omitempty"`
	Maxcwnd                     bool   `json:"maxcwnd,omitempty"`
	Fack                        bool   `json:"fack,omitempty"`
	Tcpmode                     bool   `json:"tcpmode,omitempty"`
	Tcpfastopen                 bool   `json:"tcpfastopen,omitempty"`
	Hystart                     bool   `json:"hystart,omitempty"`
	Dupackthresh                bool   `json:"dupackthresh,omitempty"`
	Burstratecontrol            bool   `json:"burstratecontrol,omitempty"`
	Tcprate                     bool   `json:"tcprate,omitempty"`
	Rateqmax                    bool   `json:"rateqmax,omitempty"`
	Drophalfclosedconnontimeout bool   `json:"drophalfclosedconnontimeout,omitempty"`
	Dropestconnontimeout        bool   `json:"dropestconnontimeout,omitempty"`
}

type update_nstcpprofile struct {
	Name                        string `json:"name"`
	Ws                          string `json:"ws,omitempty"`
	Sack                        string `json:"sack,omitempty"`
	Wsval                       int    `json:"wsval,string,omitempty"`
	Nagle                       string `json:"nagle,omitempty"`
	Ackonpush                   string `json:"ackonpush,omitempty"`
	Mss                         int    `json:"mss,string,omitempty"`
	Maxburst                    int    `json:"maxburst,string,omitempty"`
	Initialcwnd                 int    `json:"initialcwnd,string,omitempty"`
	Delayedack                  int    `json:"delayedack,string,omitempty"`
	Oooqsize                    int    `json:"oooqsize,string,omitempty"`
	Maxpktpermss                int    `json:"maxpktpermss,string,omitempty"`
	Pktperretx                  int    `json:"pktperretx,string,omitempty"`
	Minrto                      int    `json:"minrto,string,omitempty"`
	Slowstartincr               int    `json:"slowstartincr,string,omitempty"`
	Buffersize                  int    `json:"buffersize,string,omitempty"`
	Syncookie                   string `json:"syncookie,omitempty"`
	Kaprobeupdatelastactivity   string `json:"kaprobeupdatelastactivity,omitempty"`
	Flavor                      string `json:"flavor,omitempty"`
	Dynamicreceivebuffering     string `json:"dynamicreceivebuffering,omitempty"`
	Ka                          string `json:"ka,omitempty"`
	Kaconnidletime              int    `json:"kaconnidletime,string,omitempty"`
	Kamaxprobes                 int    `json:"kamaxprobes,string,omitempty"`
	Kaprobeinterval             int    `json:"kaprobeinterval,string,omitempty"`
	Sendbuffsize                int    `json:"sendbuffsize,string,omitempty"`
	Mptcp                       string `json:"mptcp,omitempty"`
	Establishclientconn         string `json:"establishclientconn,omitempty"`
	Tcpsegoffload               string `json:"tcpsegoffload,omitempty"`
	Rstwindowattenuate          string `json:"rstwindowattenuate,omitempty"`
	Rstmaxack                   string `json:"rstmaxack,omitempty"`
	Spoofsyndrop                string `json:"spoofsyndrop,omitempty"`
	Ecn                         string `json:"ecn,omitempty"`
	Mptcpdropdataonpreestsf     string `json:"mptcpdropdataonpreestsf,omitempty"`
	Mptcpfastopen               string `json:"mptcpfastopen,omitempty"`
	Mptcpsessiontimeout         int    `json:"mptcpsessiontimeout,string,omitempty"`
	Timestamp                   string `json:"timestamp,omitempty"`
	Dsack                       string `json:"dsack,omitempty"`
	Ackaggregation              string `json:"ackaggregation,omitempty"`
	Frto                        string `json:"frto,omitempty"`
	Maxcwnd                     int    `json:"maxcwnd,string,omitempty"`
	Fack                        string `json:"fack,omitempty"`
	Tcpmode                     string `json:"tcpmode,omitempty"`
	Tcpfastopen                 string `json:"tcpfastopen,omitempty"`
	Hystart                     string `json:"hystart,omitempty"`
	Dupackthresh                int    `json:"dupackthresh,string,omitempty"`
	Burstratecontrol            string `json:"burstratecontrol,omitempty"`
	Tcprate                     int    `json:"tcprate,string,omitempty"`
	Rateqmax                    int    `json:"rateqmax,string,omitempty"`
	Drophalfclosedconnontimeout string `json:"drophalfclosedconnontimeout,omitempty"`
	Dropestconnontimeout        string `json:"dropestconnontimeout,omitempty"`
}

type rename_nstcpprofile struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_nstcpprofile_payload struct {
	Resource Nstcpprofile `json:"nstcpprofile"`
}

type rename_nstcpprofile_payload struct {
	Rename rename_nstcpprofile `json:"nstcpprofile"`
}

type unset_nstcpprofile_payload struct {
	Unset NstcpprofileUnset `json:"nstcpprofile"`
}

type update_nstcpprofile_payload struct {
	Update update_nstcpprofile `json:"nstcpprofile"`
}

type get_nstcpprofile_result struct {
	Results []Nstcpprofile `json:"nstcpprofile"`
}

type count_nstcpprofile_result struct {
	Results []Count `json:"nstcpprofile"`
}

func (c *NitroClient) AddNstcpprofile(resource Nstcpprofile) error {
	payload := add_nstcpprofile_payload{
		resource,
	}

	return c.post("nstcpprofile", "", nil, payload)
}

func (c *NitroClient) RenameNstcpprofile(name string, newName string) error {
	payload := rename_nstcpprofile_payload{
		rename_nstcpprofile{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("nstcpprofile", "", qs, payload)
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

func (c *NitroClient) ExistsNstcpprofile(key string) (bool, error) {
	var results count_nstcpprofile_result

	id, qs := nstcpprofile_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("nstcpprofile", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListNstcpprofile() ([]Nstcpprofile, error) {
	results := get_nstcpprofile_result{}

	if err := c.get("nstcpprofile", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetNstcpprofile(key string) (*Nstcpprofile, error) {
	var results get_nstcpprofile_result

	id, qs := nstcpprofile_key_to_id_args(key)

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

func (c *NitroClient) DeleteNstcpprofile(key string) error {
	id, qs := nstcpprofile_key_to_id_args(key)

	return c.delete("nstcpprofile", id, qs)
}

func (c *NitroClient) UnsetNstcpprofile(unset NstcpprofileUnset) error {
	payload := unset_nstcpprofile_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("nstcpprofile", "", qs, payload)
}

func (c *NitroClient) UpdateNstcpprofile(resource Nstcpprofile) error {
	payload := update_nstcpprofile_payload{
		update_nstcpprofile{
			resource.Name,
			resource.Ws,
			resource.Sack,
			resource.Wsval,
			resource.Nagle,
			resource.Ackonpush,
			resource.Mss,
			resource.Maxburst,
			resource.Initialcwnd,
			resource.Delayedack,
			resource.Oooqsize,
			resource.Maxpktpermss,
			resource.Pktperretx,
			resource.Minrto,
			resource.Slowstartincr,
			resource.Buffersize,
			resource.Syncookie,
			resource.Kaprobeupdatelastactivity,
			resource.Flavor,
			resource.Dynamicreceivebuffering,
			resource.Ka,
			resource.Kaconnidletime,
			resource.Kamaxprobes,
			resource.Kaprobeinterval,
			resource.Sendbuffsize,
			resource.Mptcp,
			resource.Establishclientconn,
			resource.Tcpsegoffload,
			resource.Rstwindowattenuate,
			resource.Rstmaxack,
			resource.Spoofsyndrop,
			resource.Ecn,
			resource.Mptcpdropdataonpreestsf,
			resource.Mptcpfastopen,
			resource.Mptcpsessiontimeout,
			resource.Timestamp,
			resource.Dsack,
			resource.Ackaggregation,
			resource.Frto,
			resource.Maxcwnd,
			resource.Fack,
			resource.Tcpmode,
			resource.Tcpfastopen,
			resource.Hystart,
			resource.Dupackthresh,
			resource.Burstratecontrol,
			resource.Tcprate,
			resource.Rateqmax,
			resource.Drophalfclosedconnontimeout,
			resource.Dropestconnontimeout,
		},
	}

	return c.put("nstcpprofile", "", nil, payload)
}
