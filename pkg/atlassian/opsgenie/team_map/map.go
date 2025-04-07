package team_map

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/team"
	"github.com/funtimecoding/go-library/pkg/face"
)

type Map struct {
	Teams []*team.Team

	TeamMap      map[string]*team.Team
	KeyByNameMap map[string]string

	tagToName face.SliceAlias
}
