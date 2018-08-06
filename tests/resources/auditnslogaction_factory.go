package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func setup_auditnslogaction(t *testing.T, client *nitro.NitroClient) (*nitro.Auditnslogaction, func()) {
	//resource := nitro.Auditnslogaction {
	//Name: "auditnslogaction",
	//Acl: (ENABLED|DISABLED),
	//Alg: (ENABLED|DISABLED),
	//Appflowexport: (ENABLED|DISABLED),
	//Dateformat: (MMDDYYYY|DDMMYYYY|YYYYMMDD),
	//Domainresolveretry: int,
	//Logfacility: (LOCAL0|LOCAL1|LOCAL2|LOCAL3|LOCAL4|LOCAL5|LOCAL6|LOCAL7),
	//Loglevel: (ALL|EMERGENCY|ALERT|CRITICAL|ERROR|WARNING|NOTICE|INFORMATIONAL|DEBUG|NONE)[],
	//Lsn: (ENABLED|DISABLED),
	//Serverdomainname: string,
	//Serverip: ip,
	//Serverport: int,
	//Sslinterception: (ENABLED|DISABLED),
	//Subscriberlog: (ENABLED|DISABLED),
	//Tcp: (NONE|ALL),
	//Timezone: (GMT_TIME|LOCAL_TIME),
	//Userdefinedauditlog: (YES|NO),
	//}

	return nil, func() {
	}
}
