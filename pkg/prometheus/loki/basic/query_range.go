package basic

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic/constant"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic/response"
	"log"
	"net/url"
	"time"
)

func (c *Client) QueryRange(
	query string,
	start time.Time,
	end time.Time,
) response.Data {
	r := &response.Query{}
	notation.DecodeStrict(
		c.Get(
			fmt.Sprintf(
				"/query_range?start=%d&end=%d&query=%s",
				start.Unix(),
				end.Unix(),
				url.QueryEscape(query),
			),
		),
		&r,
		false,
	)

	if r.Status != constant.Success {
		log.Panicf("unexpected status: %s", r.Status)
	}

	return r.Data
}
