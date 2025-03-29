package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/schedule"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/ptr"
	raw "github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
	"sort"
)

func (c *Client) Schedules() []*schedule.Schedule {
	response, e := c.userClient.Schedule.List(
		c.context,
		&raw.ListRequest{Expand: ptr.To[bool](true)},
	)
	errors.PanicOnError(e)
	result := schedule.NewSlice(response.Schedule, c.TeamMap())
	sort.SliceStable(
		result,
		func(
			i int,
			j int,
		) bool {
			return result[i].Name < result[j].Name
		},
	)

	return result
}
