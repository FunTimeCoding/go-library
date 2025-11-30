package loki

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/message"
	"time"
)

func (c *Client) QueryRange(
	query string,
	start time.Time,
	end time.Time,
) ([]*message.Message, *message.Meta) {
	return message.NewSlice(c.basic.QueryRange(query, start, end))
}
