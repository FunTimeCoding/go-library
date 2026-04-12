package basic

import (
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic/constant"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic/response"
	"github.com/funtimecoding/go-library/pkg/web/parameter"
	"log"
	"time"
)

func (c *Client) QueryRange(
	query string,
	start time.Time,
	end time.Time,
	limit int,
) response.QueryResult {
	r := response.NewQuery()
	notation.DecodeStrict(
		c.Get(
			c.base.Copy().Path(constant.QueryRange).SetInteger64(
				parameter.Start,
				start.UnixNano(),
			).SetInteger64(
				parameter.End,
				end.UnixNano(),
			).Set(
				parameter.Query,
				query,
			).SetInteger(
				parameter.Limit,
				limit,
			).String(),
		),
		&r,
		true,
	)

	if r.Status != constant.Success {
		log.Panicf("unexpected status: %s", r.Status)
	}

	return r.Result
}
