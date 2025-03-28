package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/logs"
)

func (c *Client) Logs(marker string) []logs.Log {
	result, e := c.userClient.Log.ListLogFiles(
		c.context,
		&logs.ListLogFilesRequest{Marker: marker},
	)
	errors.PanicOnError(e)

	return result.Logs
}
