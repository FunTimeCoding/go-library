package basic

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic/constant"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic/response"
	"log"
	"time"
)

func (c *Client) LabelValues(
	start time.Time,
	end time.Time,
	label string,
) []string {
	r := &response.List{}
	notation.DecodeStrict(
		c.Get(
			fmt.Sprintf(
				"/label/%s/values?start=%d&end=%d",
				label,
				start.Unix(),
				end.Unix(),
			),
		),
		&r,
		true,
	)

	if r.Status != constant.Success {
		log.Panicf("unexpected status: %s", r.Status)
	}

	return r.Data
}
