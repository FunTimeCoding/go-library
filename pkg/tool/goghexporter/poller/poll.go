package poller

import "github.com/funtimecoding/go-library/pkg/github/run"

func (p *Poller) Poll() {
	type key struct {
		repo     string
		workflow string
		branch   string
	}
	latest := make(map[key]*run.Run)

	for _, repo := range p.client.MustRepositories(p.owner) {
		name := repo.GetName()

		for _, r := range p.client.MustLatestRuns(p.owner, name) {
			if r.Status != run.Completed {
				continue
			}

			k := key{name, r.Name, r.Branch}

			if existing, okay := latest[k]; !okay || r.Identifier > existing.Identifier {
				latest[k] = r
			}
		}
	}

	p.gauge.Reset()

	for k, r := range latest {
		p.gauge.WithLabelValues(
			p.owner,
			k.repo,
			k.workflow,
			k.branch,
			r.Conclusion,
		).Set(1)
	}
}
