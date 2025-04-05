package score_colorer

import "github.com/funtimecoding/go-library/pkg/face"

func GreenRed[T face.ScoreColorable](v ...T) *Colorer {
	return Dynamic("#00aa00", "#aa0000", v...)
}
