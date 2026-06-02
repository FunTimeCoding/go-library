package alertmanager

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/silence"

func (c *Client) SilenceByRule(name string) (*silence.Silence, error) {
	v, e := c.Silences(false)

	if e != nil {
		return nil, e
	}

	for _, s := range v {
		if s.Rule == name {
			return s, nil
		}
	}

	return nil, nil
}
