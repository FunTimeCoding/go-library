package alertmanager

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/silence"

func (c *Client) SilenceByRule(name string) *silence.Silence {
	for _, s := range c.Silences(false) {
		if s.Rule == name {
			return s
		}
	}

	return nil
}
