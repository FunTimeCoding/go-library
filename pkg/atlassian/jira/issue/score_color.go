package issue

import "github.com/funtimecoding/go-library/pkg/face"

func (i *Issue) ScoreColor() face.SprintFunction {
	return i.scoreColor
}
