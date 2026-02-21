package gomaintlog

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/maintenance_log"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	monitor.ParseBind(version, gitHash, buildDate)
	s := server.New(maintenance_log.New().Nested())
	m := http.NewServeMux()
	s.Setup(m)
	h := web.Server(m, server.Address)
	web.ServeAsynchronous(h)
	system.KillSignalBlock()
	web.GracefulShutdown(context.Background(), h, true)
}
