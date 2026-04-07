package tool

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/monitor"
)

type Tool struct {
	client  *mattermost.Client
	monitor *monitor.Monitor
}
