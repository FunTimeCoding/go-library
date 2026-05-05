package gomonitord

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/monitor/coder"
	"github.com/funtimecoding/go-library/pkg/monitor/gorilla"
	"github.com/funtimecoding/go-library/pkg/monitor/gorilla/example_client"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/constant"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New("gomonitord", version)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	monitor.ParseBind(version, gitHash, buildDate)

	if false {
		// Start server: go run cmd/gomonitord/main.go localhost:3002
		// Connect: websocat ws://127.0.0.1:3002 --protocol echo
		errors.PanicOnError(coder.Run())
	}

	if false {
		example_client.Run()
	}

	if true {
		gorilla.Run(
			web.AddressHostPort(constant.Localhost, constant.ListenPort),
		)
	}
}
