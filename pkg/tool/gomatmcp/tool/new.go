package tool

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/monitor"
)

func New(
	m *mattermost.Client,
	o *monitor.Monitor,
) *Tool {
	return &Tool{client: m, monitor: o}
}
