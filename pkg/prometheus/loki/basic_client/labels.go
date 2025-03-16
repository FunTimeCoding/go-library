package basic_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic_client/constant"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic_client/response"
	"log"
	"time"
)

func (c *Client) Labels(
	start time.Time,
	end time.Time,
) []string {
	result := &response.Labels{}
	notation.DecodeStrict(
		c.Get(
			fmt.Sprintf(
				"/labels?start=%d&end=%d",
				start.Unix(),
				end.Unix(),
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
