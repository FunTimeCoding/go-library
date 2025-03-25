package score_colorer

import "github.com/funtimecoding/go-library/pkg/math/range_mapping"

func (c *Colorer) Mapping() []*range_mapping.Mapping {
	return c.mapping
}
