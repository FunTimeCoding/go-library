package option

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/team_map"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/user_map"
)

type Alert struct {
	Team    *team_map.Map
	User    *user_map.Map
	WebHost string

	ShortAlert        constant.StringAlias
	ShortUser         constant.StringAlias
	DescriptionToName constant.StringAlias
	ParseDescription  constant.ParseDescription
}
