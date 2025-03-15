package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/go-openapi/strfmt"
	"github.com/prometheus/alertmanager/api/v2/client/silence"
)

func (c *Client) DeleteSilence(identifier string) {
	// Does not delete: https://github.com/prometheus/alertmanager/issues/3614
	p := silence.NewDeleteSilenceParams()
	p.SilenceID = strfmt.UUID(identifier)
	_, e := c.client.Silence.DeleteSilence(p)
	errors.PanicOnError(e)
}
