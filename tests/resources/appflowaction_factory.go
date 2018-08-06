package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func setup_appflowaction(t *testing.T, client *nitro.NitroClient) (*nitro.Appflowaction, func()) {
	collector, collectorTearDown := setup_appflowcollector(t, client)

	client.AddAppflowcollector(*collector)

	resource := nitro.Appflowaction{
		Name: "appflowaction",
		Clientsidemeasurements: "ENABLED",
		Comment:                "comment",
		Pagetracking:           "ENABLED",
		Securityinsight:        "ENABLED",
		Webinsight:             "ENABLED",
	}

	resource.Collectors = append(resource.Collectors, collector.Name)

	return &resource, func() {
		client.DeleteAppflowcollector(collector.Name)

		if collectorTearDown != nil {
			collectorTearDown()
		}
	}
}
