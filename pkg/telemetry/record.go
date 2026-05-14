package telemetry

import "github.com/funtimecoding/go-library/pkg/telemetry/record"

func (c *Client) Record(r *record.Record) {
	go c.send(r)
}
