package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func Setup_appflowpolicy(t *testing.T, client *nitro.NitroClient) (*nitro.Appflowpolicy, func()) {
	action, actionTearDown := Setup_appflowaction(t, client)

	client.AddAppflowaction(*action)

	resource := nitro.Appflowpolicy{
		Name:        "appflowpolicy",
		Action:      action.Name,
		Comment:     "comment",
		Rule:        "TRUE",
		Undefaction: action.Name,
	}

	return &resource, func() {
		client.DeleteAppflowaction(action.ToKey())

		if actionTearDown != nil {
			actionTearDown()
		}
	}
}
