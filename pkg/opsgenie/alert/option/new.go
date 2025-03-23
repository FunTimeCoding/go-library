package option

import (
	"github.com/funtimecoding/go-library/pkg/opsgenie/constant"
	"github.com/funtimecoding/go-library/pkg/opsgenie/team_map"
	"github.com/funtimecoding/go-library/pkg/opsgenie/user_map"
)

func New(
	t *team_map.Map,
	u *user_map.Map,
	webHost string,
	alert constant.StringAlias,
	user constant.StringAlias,
	description constant.StringAlias,
) *Alert {
	return &Alert{
		Team:              t,
		User:              u,
		WebHost:           webHost,
		ShortAlert:        alert,
		ShortUser:         user,
		DescriptionToName: description,
	}
}
