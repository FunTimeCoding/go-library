package mock_client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/habitica/request"
	"github.com/funtimecoding/go-library/pkg/habitica/response"
)

func (c *Client) MustCreateTask(b request.CreateTaskBody) response.Task {
	result, e := c.CreateTask(b)
	errors.PanicOnError(e)

	return result
}
