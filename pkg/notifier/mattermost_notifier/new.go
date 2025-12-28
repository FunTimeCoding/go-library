package mattermost_notifier

import "github.com/funtimecoding/go-library/pkg/chat/mattermost"

func New(
	m *mattermost.Client,
	channel string,
	prefix string,
) *Notifier {
	h := m.TeamChannel(channel)

	return &Notifier{mattermost: m, channel: h, prefix: prefix}
}
