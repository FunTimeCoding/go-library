package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/team"
	"github.com/funtimecoding/go-library/pkg/errors"
	raw "github.com/opsgenie/opsgenie-go-sdk-v2/team"
)

func (c *Client) Teams() []*team.Team {
	result, e := c.userClient.Team.List(c.context, &raw.ListTeamRequest{})
	errors.PanicOnError(e)

	return team.NewSlice(result.Teams)
}
