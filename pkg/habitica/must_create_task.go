package habitica

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/habitica/request"
	"github.com/funtimecoding/go-library/pkg/habitica/response"
)

func (c *Client) MustCreateTask(body request.CreateTaskBody) response.Task {
	result, e := c.CreateTask(body)
	errors.PanicOnError(e)

	return result
}
