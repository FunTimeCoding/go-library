package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/mattermost/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment(o ...OptionFunc) *Client {
	if s := environment.GetDefault(
		constant.HostEnvironment,
		"",
	); s != "" {
		o = append(o, WithHost(s))
	}

	if s := environment.GetDefault(
		constant.TokenEnvironment,
		"",
	); s != "" {
		o = append(o, WithToken(s))
	}

	if s := environment.GetDefault(
		constant.TeamEnvironment,
		"",
	); s != "" {
		o = append(o, WithTeam(s))
	}

	if s := environment.GetDefault(
		constant.ChannelEnvironment,
		"",
	); s != "" {
		o = append(o, WithChannel(s))
	}

	return New(o...)
}
