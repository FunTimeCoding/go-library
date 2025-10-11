package loki

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
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

	return &Client{basic: basic.New(host, user, password)}
}

func slogWay(host string) {
	// Another way to log, not to read
	configuration, e := loki.NewDefaultConfig(
		locator.New(host).Base(constant.Base).Path(constant.Push).String(),
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
