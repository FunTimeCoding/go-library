package monitor

import (
	"github.com/funtimecoding/go-library/pkg/bubbletea/table/item"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/gorilla/client"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func New(connect bool) *Model {
	auto := true

	if s := environment.Default(
		constant.ManualEnvironment,
		"",
	); s != "" {
		auto = false
	}

	return &Model{
		table:    item.New(connect),
		client:   client.New(),
		connect:  connect,
		user:     system.User().Username,
		hostname: system.Hostname(),
		auto:     auto,
	}
}
