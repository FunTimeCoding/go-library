package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/rotation"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/schedule"
	"github.com/funtimecoding/go-library/pkg/errors"
	rawSchedule "github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
)

func (c *Client) ScheduleRotations(v *schedule.Schedule) []*rotation.Rotation {
	result, e := c.userClient.Schedule.ListRotations(
		c.context,
		&rawSchedule.ListRotationsRequest{
			ScheduleIdentifierType:  rawSchedule.Id,
			ScheduleIdentifierValue: v.Identifier,
		},
	)
	errors.PanicOnError(e)

	return rotation.NewSlice(result.Rotations)
}
