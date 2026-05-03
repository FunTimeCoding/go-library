package habitica

import (
	"github.com/funtimecoding/go-library/pkg/habitica/response"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (c *Client) Allocate(stat string) (response.Stats, error) {
	var result response.Stats

	return result, c.post(
		join.Empty("/user/allocate?stat=", stat),
		nil,
		&result,
	)
}
