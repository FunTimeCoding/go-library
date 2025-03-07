package console

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/math/range_mapping"
)

func ScoreMappings(v []face.ScoreColorable) []*range_mapping.Mapping {
	return range_mapping.Generate(
		CollectScores(v),
		[]string{GreenColor, YellowColor, RedColor},
	)
}
