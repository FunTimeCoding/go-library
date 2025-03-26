package age_colorer

import "github.com/funtimecoding/go-library/pkg/face"

func GreenRed(v ...face.AgeColorable) *Colorer {
	return Dynamic("#00ff00", "#ff0000", v...)
}
