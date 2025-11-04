package kestra

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/kestra/constant"
	"github.com/kestra-io/client-sdk/go-sdk"
)

func (c *Client) Logs(execution string) []kestra_api_client.LogEntry {
	result, r, e := c.client.LogsAPI.ListLogsFromExecution(
		c.context,
		execution,
		constant.MainTenant,
	).Execute()
	errors.PanicOnWebError(r, e)

	return result
}
