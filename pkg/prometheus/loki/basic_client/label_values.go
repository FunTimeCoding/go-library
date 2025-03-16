package basic_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic_client/constant"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic_client/response"
	"log"
	"time"
)

func (c *Client) LabelValues(
	start time.Time,
	end time.Time,
	label string,
) []string {
	result := &response.List{}
	notation.DecodeStrict(
		c.Get(
			fmt.Sprintf(
				"/label/%s/values?start=%d&end=%d",
				label,
				start.Unix(),
				end.Unix(),
			),
		),
		&result,
		true,
	)

	if result.Status != constant.Success {
		log.Panicf("unexpected status: %s", result.Status)
	}

	return result.Data
}
