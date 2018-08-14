package resources

import (
	"testing"

	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
)

func Setup_service(t *testing.T, client *nitro.NitroClient) (*nitro.Service, func()) {
	//dnsprofile, dnsprofileTearDown := Setup_dnsprofile(t, client)
	//nshttpprofile, nshttpprofileTearDown := Setup_nshttpprofile(t, client)

	// resource := nitro.Service{
	// 	Name:          "service",
	// 	Accessdown:    "YES",
	// 	Appflowlog:    "DISABLED",
	// 	Cacheable:     "YES",
	// 	Cachetype:     "TRANSPARENT",
	// 	Cip:           "ENABLED",
	// 	Cipheader:     "header",
	// 	Cka:           "YES",
	// 	Cleartextport: 1290,
	// 	Clttimeout:    60,
	// 	Cmp:           "NO",
	// 	Comment:       "Comment service",
	// 	//Customserverid:     string,
	// 	//Dnsprofilename:     dnsprofile.Dnsprofilename,
	// 	//Downstateflush:     "ENABLED",
	// 	//Hashid:             10,
	// 	Healthmonitor:      "YES",
	// 	Httpprofilename:    nshttpprofile.Name,
	// 	Ip:                 "192.10.2.1",
	// 	Maxbandwidth:       8096,
	// 	Maxclient:          100,
	// 	Maxreq:             100,
	// 	Monconnectionclose: "RESET",
	// 	Monthreshold:       20,
	// 	Netprofile:         "",
	// 	Pathmonitor:        "YES",
	// 	Pathmonitorindv:    "YES",
	// 	Port:               80,
	// 	Processlocal:       "ENABLED",
	// 	Rtspsessionidremap: "ON",
	// 	Sc:                 "OFF",
	// 	// Serverid:           double,
	// 	// Servername:         server.name,
	// 	Servicetype: "HTTP",
	// 	// Sp:                 "OFF",
	// 	// Svrtimeout:         60,
	// 	// Tcpb:               "NO",
	// 	// Tcpprofilename:     nstcpprofile.name,
	// 	// Td:                 double,
	// 	Useproxyport: "NO",
	// 	// Usip:               (YES | NO),
	// }

	return nil, func() {
		//client.DeleteService(resource.ToKey())
		// if dnsprofileTearDown != nil {
		// 	dnsprofileTearDown()
		// }
		// if nshttpprofileTearDown != nil {
		// 	nshttpprofileTearDown()
		// }
	}
}
