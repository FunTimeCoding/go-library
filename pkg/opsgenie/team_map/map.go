package team_map

import (
	"github.com/funtimecoding/go-library/pkg/opsgenie/constant"
	"github.com/funtimecoding/go-library/pkg/opsgenie/team"
)

type Map struct {
	Teams []*team.Team

	TeamMap      map[string]*team.Team
	KeyByNameMap map[string]string

	tagsToName constant.SliceAlias
}
