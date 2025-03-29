package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
)

func (c *Client) OnCalls(scheduleName string) *schedule.GetOnCallsResult {
	result, e := c.userClient.Schedule.GetOnCalls(
		c.context,
		&schedule.GetOnCallsRequest{
			ScheduleIdentifier: scheduleName,
			Flat:               ptr.To[bool](true),
		},
	)
	errors.PanicOnError(e)

	return result
}
