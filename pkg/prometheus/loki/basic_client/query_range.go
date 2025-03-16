package basic_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic_client/constant"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic_client/response"
	"log"
	"net/url"
	"time"
)

func (c *Client) QueryRange(
	query string,
	start time.Time,
	end time.Time,
) response.Data {
	result := &response.Query{}
	notation.DecodeStrict(
		c.Get(
			fmt.Sprintf(
				"/query_range?start=%d&end=%d&query=%s",
				start.Unix(),
				end.Unix(),
				url.QueryEscape(query),
			),
		),
		&result,
		false,
	)

	if result.Status != constant.Success {
		log.Panicf("unexpected status: %s", result.Status)
	}

	return result.Data
}
