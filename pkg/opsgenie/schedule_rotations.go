package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/opsgenie/rotation"
	"github.com/funtimecoding/go-library/pkg/opsgenie/schedule"
	raw "github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
)

func (c *Client) ScheduleRotations(v *schedule.Schedule) []*rotation.Rotation {
	result, e := c.userClient.Schedule.ListRotations(
		c.context,
		&raw.ListRotationsRequest{
			ScheduleIdentifierType:  raw.Id,
			ScheduleIdentifierValue: v.Identifier,
		},
	)
	errors.PanicOnError(e)

	return rotation.NewSlice(result.Rotations)
}
