package poller

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/getsentry/sentry-go"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	client *github.Client,
	owner string,
	interval time.Duration,
	r *prometheus.Registry,
	l *logger.Logger,
	h *sentry.Hub,
) *Poller {
	g := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "github_workflow_run_status",
			Help: "Latest completed workflow run per repo/workflow/branch (1 per entry)",
		},
		[]string{"owner", "repo", "workflow", "branch", "conclusion"},
	)
	r.MustRegister(g)

	return &Poller{
		client:   client,
		owner:    owner,
		interval: interval,
		gauge:    g,
		recovery: recovery.New(l, h),
		stop:     make(chan struct{}),
	}
}
