package basic

import (
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic/constant"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic/response"
	"github.com/funtimecoding/go-library/pkg/web/parameter"
	"log"
	"time"
)

func (c *Client) Query(query string) response.Data {
	r := &response.Query{}
	notation.DecodeStrict(
		c.Get(
			c.base.Copy().Path(constant.Query).SetInteger64(
				parameter.Time,
				time.Now().Unix(),
			).Set(parameter.Query, query).String(),
		),
		&r,
		false,
	)

	if r.Status != constant.Success {
		log.Panicf("unexpected status: %s", r.Status)
	}

	return r.Data
}
