package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/mattermost/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Get(constant.HostEnvironment, 1),
		environment.Get(constant.TokenEnvironment, 1),
		environment.GetDefault(constant.TeamEnvironment, ""),
		environment.GetDefault(constant.ChannelEnvironment, ""),
	)
}
