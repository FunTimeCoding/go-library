package poller

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
)

func (p *Poller) RecoverStale() {
	defer func() {
		if v := recover(); v != nil {
			p.logger.Structured(
				"recover stale failed",
				"error", fmt.Sprint(v),
			)
		}
	}()

	alerts, _ := p.client.Alerts(&advanced_option.Alert{})
	current := make(map[string]bool)

	for _, a := range alerts {
		if a.State == constant.ActiveState {
			current[a.Fingerprint] = true
		}
	}

	unresolved := p.store.Unresolved()

	for _, u := range unresolved {
		if current[u.Fingerprint] {
			p.firing[u.Fingerprint] = u.Key
		}
	}

	var resolved int

	for _, u := range unresolved {
		if p.firing[u.Fingerprint] == u.Key {
			continue
		}

		p.store.Resolve(u.Key)
		resolved++
	}

	p.logger.Structured(
		"recovered stale records",
		"adopted", len(p.firing),
		"resolved", resolved,
	)
}
