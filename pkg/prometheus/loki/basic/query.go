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

func (c *Client) Query(query string) response.Data {
	r := &response.Query{}
	notation.DecodeStrict(
		c.Get(
			fmt.Sprintf(
				"/query?time=%d&query=%s",
				time.Now().Unix(),
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
