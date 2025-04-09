package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/override"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/schedule"
	"github.com/funtimecoding/go-library/pkg/errors"
	rawSchedule "github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
)

func (c *Client) ScheduleOverrides(v *schedule.Schedule) []*override.Override {
	response, e := c.userClient.Schedule.ListScheduleOverride(
		c.context,
		&rawSchedule.ListScheduleOverrideRequest{
			ScheduleIdentifierType: rawSchedule.Id,
			ScheduleIdentifier:     v.Identifier,
		},
	)
	errors.PanicOnError(e)

	return override.NewSlice(response.ScheduleOverride)
}
