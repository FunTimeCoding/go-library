package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/incident"
)

func (c *Client) Incidents(query string) []incident.Incident {
	result, e := c.userClient.Incident.List(
		c.context,
		&incident.ListRequest{Query: query},
	)
	errors.PanicOnError(e)

	return result.Incidents
}
