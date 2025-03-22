package console

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/math/range_mapping"
)

func ScoreMapping(v []face.ScoreColorable) []*range_mapping.Mapping {
	return range_mapping.Generate(
		CollectScore(v),
		[]string{GreenColor, YellowColor, RedColor},
	)
}
