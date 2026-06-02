package alertmanager

import (
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/prometheus/alertmanager/api/v2/client/silence"
)

func (c *Client) DeleteSilence(identifier string) error {
	p := silence.NewDeleteSilenceParams()
	p.SilenceID = strfmt.UUID(identifier)
	result, e := c.client.Silence.DeleteSilence(p)

	if e != nil {
		return e
	}

	if !result.IsSuccess() {
		return fmt.Errorf("unexpected status code: %+v", result)
	}

	return nil
}
