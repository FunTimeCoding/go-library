package poller

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlog/store"
	"time"
)

func (p *Poller) Poll() {
	p.lastPoll.Store(time.Now())

	defer func() {
		if v := recover(); v != nil {
			p.logger.Structured(
				"poll failed",
				"error", fmt.Sprint(v),
			)
		}
	}()

	alerts, _ := p.client.Alerts(&advanced_option.Alert{})
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

		p.firing[a.Fingerprint] = p.store.Save(r)
	}

	for fingerprint, key := range p.firing {
		if current[fingerprint] {
			continue
		}

		p.store.Resolve(key)
		delete(p.firing, fingerprint)
	}

	if pruned := p.store.Prune(time.Now().Add(-p.retention)); pruned > 0 {
		p.logger.Structured("pruned records", "count", pruned)
	}
}
