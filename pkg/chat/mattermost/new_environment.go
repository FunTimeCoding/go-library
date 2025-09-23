package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment(o ...Option) *Client {
	if s := environment.Fallback(
		constant.HostEnvironment,
		"",
	); s != "" {
		o = append(o, WithHost(s))
	}

	if s := environment.Fallback(
		constant.TokenEnvironment,
		"",
	); s != "" {
		o = append(o, WithToken(s))
	}

	if s := environment.Fallback(
		constant.TeamEnvironment,
		"",
	); s != "" {
		o = append(o, WithTeam(s))
	}

	if s := environment.Fallback(
		constant.ChannelEnvironment,
		"",
	); s != "" {
		o = append(o, WithChannel(s))
	}

	return New(o...)
}
