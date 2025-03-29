package team_map

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/team"
)

type Map struct {
	Teams []*team.Team

	TeamMap      map[string]*team.Team
	KeyByNameMap map[string]string

	tagToName constant.SliceAlias
}
