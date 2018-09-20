package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func Setup_appflowaction(t *testing.T, client *nitro.NitroClient) (*nitro.Appflowaction, func()) {
	collector, collectorTearDown := Setup_appflowcollector(t, client)

	client.AddAppflowcollector(*collector)

	resource := nitro.Appflowaction{
		Name:                   "appflowaction",
		Clientsidemeasurements: "ENABLED",
		Comment:                "comment",
		Pagetracking:           "ENABLED",
		Securityinsight:        "ENABLED",
		Webinsight:             "ENABLED",
	}

	resource.Collectors = append(resource.Collectors, collector.Name)

	return &resource, func() {
		client.DeleteAppflowcollector(collector.ToKey())

		if collectorTearDown != nil {
			collectorTearDown()
		}
	}
}
