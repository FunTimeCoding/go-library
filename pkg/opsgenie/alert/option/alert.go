package option

import (
	"github.com/funtimecoding/go-library/pkg/opsgenie/constant"
	"github.com/funtimecoding/go-library/pkg/opsgenie/team_map"
	"github.com/funtimecoding/go-library/pkg/opsgenie/user_map"
)

type Alert struct {
	Team    *team_map.Map
	User    *user_map.Map
	WebHost string

	ShortAlert        constant.StringAlias
	ShortUser         constant.StringAlias
	DescriptionToName constant.StringAlias
}
