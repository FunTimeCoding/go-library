package console

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/math/range_mapping"
)

func AgeMapping(v []face.AgeColorable) []*range_mapping.Mapping {
	return range_mapping.Generate(
		CollectAge(v),
		[]string{GreenColor, YellowColor, RedColor},
	)
}
