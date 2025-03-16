package loki

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic_client"
	"github.com/grafana/loki-client-go/loki"
	"github.com/samber/slog-loki/v3"
	"log/slog"
)

// Conflict with github.com/prometheus/common/config
//import "github.com/grafana/loki/v3/integration/client"

func New(
	host string,
	user string,
	password string,
) *Client {
	//client.New()

	if false {
		slogWay()
	}

	return &Client{basic: basic_client.New(host, user, password)}
}

func slogWay() {
	// Yet another way to log, not to read logs
	config, _ := loki.NewDefaultConfig(
		"http://localhost:3100/loki/api/v1/push",
	)
	config.TenantID = "xyz"
	client, _ := loki.New(config)
	l := slog.New(
		slogloki.Option{
			Level:  slog.LevelDebug,
			Client: client,
		}.NewLokiHandler(),
	)
	l = l.With("environment", "dev").With("release", "v1.0.0")
	l.Error("caramba!")
	l.Info("user registration")
	client.Stop()
}
