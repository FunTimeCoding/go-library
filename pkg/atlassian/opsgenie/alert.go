package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/team"
	"github.com/funtimecoding/go-library/pkg/errors"
	raw "github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

func (c *Client) Alert(
	identifier string,
	t *team.Team,
) *alert.Alert {
	result, e := c.userClient.Alert.Get(
		c.context,
		&raw.GetAlertRequest{
			IdentifierType:  raw.ALERTID,
			IdentifierValue: identifier,
		},
	)
	errors.PanicOnError(e)

	return alert.NewDetail(result, c.AlertOption(), t)
}
