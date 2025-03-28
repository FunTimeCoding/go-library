package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/opsgenie/override"
	"github.com/funtimecoding/go-library/pkg/opsgenie/schedule"
	raw "github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
)

func (c *Client) ScheduleOverrides(v *schedule.Schedule) []*override.Override {
	response, e := c.userClient.Schedule.ListScheduleOverride(
		c.context,
		&raw.ListScheduleOverrideRequest{
			ScheduleIdentifierType: raw.Id,
			ScheduleIdentifier:     v.Identifier,
		},
	)
	errors.PanicOnError(e)

	return override.NewSlice(response.ScheduleOverride)
}
