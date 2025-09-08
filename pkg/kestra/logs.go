package kestra

import (
	"github.com/funtimecoding/go-library/pkg/kestra/constant"
	"github.com/kestra-io/client-sdk/go-sdk"
)

func (c *Client) Logs(execution string) []kestra_api_client.LogEntry {
	result, r, e := c.client.LogsAPI.ListLogsFromExecution(
		c.context,
		execution,
		constant.MainTenant,
	).Execute()
	panicOnError(e, r)

	return result
}
