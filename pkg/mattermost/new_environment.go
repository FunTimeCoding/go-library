package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/mattermost/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment(o ...OptionFunc) *Client {
	t := &Option{}

	for _, p := range o {
		p(t)
	}

	if t.Host == "" {
		t.Host = environment.Get(constant.HostEnvironment, 1)
	}

	if t.Token == "" {
		t.Token = environment.Get(constant.TokenEnvironment, 1)
	}

	if t.Team == "" {
		t.Team = environment.GetDefault(constant.TeamEnvironment, "")
	}

	if t.Channel == "" {
		t.Channel = environment.GetDefault(
			constant.ChannelEnvironment,
			"",
		)
	}

	return New(t.Host, t.Token, t.Team, t.Channel)
}
