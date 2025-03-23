package team_map

import "github.com/funtimecoding/go-library/pkg/opsgenie/constant"

func (m *Map) TagsToName(f constant.SliceAlias) {
	m.tagsToName = f
}
