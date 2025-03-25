package merge_request

import (
	"github.com/funtimecoding/go-library/pkg/console/age_colorer"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/math/range_mapping"
)

func AgeMapping(v []*Request) []*range_mapping.Mapping {
	var a []face.AgeColorable

	for _, element := range v {
		a = append(a, element)
	}

	return age_colorer.Default(a...).Mapping()
}
