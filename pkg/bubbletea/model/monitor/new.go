package monitor

import (
	"github.com/funtimecoding/go-library/pkg/bubbletea/table/item"
	"github.com/funtimecoding/go-library/pkg/monitor/gorilla/client"
	"github.com/funtimecoding/go-library/pkg/system"
)

func New(connect bool) *Model {
	return &Model{
		table:    item.New(),
		client:   client.New(),
		connect:  connect,
		user:     system.User().Username,
		hostname: system.Hostname(),
	}
}
