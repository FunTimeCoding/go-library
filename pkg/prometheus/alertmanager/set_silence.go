package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"log"
	"time"
)

func (c *Client) SetSilence(
	alert string,
	comment string,
	d time.Duration,
) string {
	if alert == "" {
		log.Panicf("alert empty")
	}

	if !c.RuleExists(alert) {
		log.Panicf("rule not found: %s", alert)
	}

	t := time.Now()

	if d == 0 {
		d = constant.DefaultDuration
	}

	if s := c.SilenceByRule(alert); s != nil {
		return c.PostSilence(
			s.Identifier,
			alert,
			comment,
			*s.Start,
			t.Add(d),
		)
	}

	return c.PostSilence("", alert, comment, t, t.Add(d))
}
