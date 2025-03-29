package team_map

import "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"

func (m *Map) KeyByName(name string) string {
	if v, okay := m.KeyByNameMap[name]; okay {
		return v
	}

	return constant.NoKey
}
