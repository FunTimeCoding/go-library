package option

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/team_map"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/user_map"
)

func New(
	t *team_map.Map,
	u *user_map.Map,
	webHost string,
	alert constant.StringAlias,
	user constant.StringAlias,
	descriptionToName constant.StringAlias,
	parseDescription constant.ParseDescription,
) *Alert {
	return &Alert{
		Team:              t,
		User:              u,
		WebHost:           webHost,
		ShortAlert:        alert,
		ShortUser:         user,
		DescriptionToName: descriptionToName,
		ParseDescription:  parseDescription,
	}
}
