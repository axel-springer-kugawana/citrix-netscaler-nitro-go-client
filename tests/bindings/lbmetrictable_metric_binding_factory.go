package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-nitro-go-client/tests/resources"
	"testing"
)

func Setup_lbmetrictable_metric_binding(t *testing.T, client *nitro.NitroClient) (*nitro.LbmetrictableMetricBinding, func()) {
	metrictable, metrictableTearDown := resources.Setup_lbmetrictable(t, client)

	client.AddLbmetrictable(*metrictable)

	resource := nitro.LbmetrictableMetricBinding{
		Metrictable: metrictable.Metrictable,
		Metric:      "metric",
		Snmpoid:     "1.3.6.1.4.1.5951.4.1.1.46.2.0",
	}

	return &resource, func() {
		client.DeleteLbmetrictable(metrictable.ToKey())

		if metrictableTearDown != nil {
			metrictableTearDown()
		}
	}
}
