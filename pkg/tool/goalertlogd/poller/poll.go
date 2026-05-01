package poller

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	"time"
)

func (p *Poller) Poll() {
	start := time.Now()
	p.lastPoll.Store(start)

	if p.metrics != nil {
		p.metrics.lastPollTime.SetToCurrentTime()
	}

	defer func() {
		if v := recover(); v != nil {
			p.logger.Structured(
				"poll failed",
				"error",
				fmt.Sprint(v),
			)
		}

		if p.metrics != nil {
			p.metrics.pollDuration.Observe(
				time.Since(start).Seconds(),
			)
		}
	}()
	alerts, _ := p.client.Alerts(advanced_option.New())
	current := make(map[string]bool)

	for _, a := range alerts {
		if a.State != constant.ActiveState {
			continue
		}

		current[a.Fingerprint] = true

		if _, exists := p.firing[a.Fingerprint]; exists {
			continue
		}

		labels := make(map[string]string)

		for k, v := range a.Labels {
			labels[k] = v
		}

		r := store.Record{
			Fingerprint: a.Fingerprint,
			Name:        a.Name,
			Severity:    a.Severity,
			Summary:     a.Summary,
			Labels:      labels,
			Start:       time.Now(),
		}
		p.firing[a.Fingerprint] = p.store.MustSave(r)

		if p.metrics != nil {
			p.metrics.alertsTotal.Inc()
		}
	}

	for fingerprint, key := range p.firing {
		if current[fingerprint] {
			continue
		}

		p.store.MustResolve(key)
		delete(p.firing, fingerprint)
	}

	if pruned := p.store.MustPrune(time.Now().Add(-p.retention)); pruned > 0 {
		p.logger.Structured("pruned records", "count", pruned)
	}

	if p.metrics != nil {
		p.metrics.alertsFiring.Set(float64(len(p.firing)))
		p.metrics.recordsTotal.Set(float64(p.store.MustCount()))
	}
}
