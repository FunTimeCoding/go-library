package basic

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic/constant"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic/response"
	"log"
	"time"
)

func (c *Client) Labels(
	start time.Time,
	end time.Time,
) []string {
	r := &response.List{}
	notation.DecodeStrict(
		c.Get(
			fmt.Sprintf(
				"/labels?start=%d&end=%d",
				start.Unix(),
				end.Unix(),
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
