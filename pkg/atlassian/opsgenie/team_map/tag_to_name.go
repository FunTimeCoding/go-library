package team_map

import "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"

func (m *Map) TagToName(f constant.SliceAlias) {
	m.tagToName = f
}
