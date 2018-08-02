package nitro

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

type NstcpprofileKey struct {
	Name string
}

type nstcpprofile_update struct {
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

type nstcpprofile_payload struct {
	nstcpprofile interface{}
}

func nstcpprofile_key_to_args(key NstcpprofileKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteNstcpprofile(key NstcpprofileKey) error {
	return c.deleteResourceWithArgs("nstcpprofile", key.Name, nstcpprofile_key_to_args(key))
}

func (c *NitroClient) GetNstcpprofile(key NstcpprofileKey) (*Nstcpprofile, error) {
	var results struct {
		Nstcpprofile []Nstcpprofile
	}

	if err := c.getResourceWithArgs("nstcpprofile", key.Name, nstcpprofile_key_to_args(key), &results); err != nil || len(results.Nstcpprofile) != 1 {
		return nil, err
	}

	return &results.Nstcpprofile[0], nil
}

func (c *NitroClient) ListNstcpprofile() ([]Nstcpprofile, error) {
	var results struct {
		Nstcpprofile []Nstcpprofile
	}

	if err := c.listResources("nstcpprofile", &results); err != nil {
		return nil, err
	}

	return results.Nstcpprofile, nil
}

func (c *NitroClient) AddNstcpprofile(resource Nstcpprofile) error {
	return c.addResource("nstcpprofile", resource)
}

func (c *NitroClient) RenameNstcpprofile(name string, newName string) error {
	return c.renameResource("nstcpprofile", "name", name, newName)
}

func (c *NitroClient) UnsetNstcpprofile(name string, fields ...string) error {
	return c.unsetResource("nstcpprofile", "name", name, fields)
}

func (c *NitroClient) UpdateNstcpprofile(resource Nstcpprofile) error {
	update := nstcpprofile_update{
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
	}

	return c.Put("nstcpprofile", update)
}
