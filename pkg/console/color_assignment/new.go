package color_assignment

import "github.com/funtimecoding/go-library/pkg/face"

func New(
	name string,
	f face.SprintFunction,
) *Assignment {
	return &Assignment{Name: name, Function: f}
}
