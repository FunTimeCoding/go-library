package loki

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic_client"
	"github.com/grafana/loki-client-go/loki"
	"github.com/samber/slog-loki/v3"
	"log/slog"
)

// Conflicts with github.com/prometheus/common/config
//import "github.com/grafana/loki/v3/integration/client"

func New(
	host string,
	user string,
	password string,
) *Client {
	//client.New()

	if false {
		slogWay(host)
	}

	return &Client{basic: basic_client.New(host, user, password)}
}

func slogWay(host string) {
	// Another way to log, not to read
	configuration, e := loki.NewDefaultConfig(
		fmt.Sprintf("https://%s/loki/api/v1/push", host),
	)
	errors.PanicOnError(e)
	configuration.TenantID = "exampleTenant"
	c, _ := loki.New(configuration)
	l := slog.New(
		slogloki.Option{Level: slog.LevelDebug, Client: c}.NewLokiHandler(),
	)
	l = l.With("environment", "dev").With("release", "v1.0.0")
	l.Error("example error")
	l.Info("example info")
	c.Stop()
}
