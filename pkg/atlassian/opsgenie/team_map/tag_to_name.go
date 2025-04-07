package team_map

import "github.com/funtimecoding/go-library/pkg/face"

func (m *Map) TagToName(f face.SliceAlias) {
	m.tagToName = f
}
