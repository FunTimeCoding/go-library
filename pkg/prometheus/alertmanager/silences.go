package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/silence"
	rawSilence "github.com/prometheus/alertmanager/api/v2/client/silence"
)

func (c *Client) Silences(expired bool) []*silence.Silence {
	p := rawSilence.NewGetSilencesParams()

	if false {
		// Does not work: https://github.com/prometheus/alertmanager/issues/2928
		if !expired {
			p.Filter = []string{"state!=expired"}
		}
	}

	response, e := c.client.Silence.GetSilences(p)
	errors.PanicOnError(e)
	var result []*silence.Silence

	for _, s := range silence.NewSlice(response.GetPayload(), c.host) {
		if !expired && s.Expired() {
			continue
		}

		result = append(result, s)
	}

	return silence.Sort(result)
}
