package branch

import (
	"github.com/funtimecoding/go-library/pkg/console/age_colorer"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/math/range_mapping"
)

func AgeMapping(v []*Branch) []*range_mapping.Mapping {
	var a []face.AgeColorable

	for _, e := range v {
		a = append(a, e)
	}

	return age_colorer.Default(a...).Mapping()
}
