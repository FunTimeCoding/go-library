package branch

import "github.com/funtimecoding/go-library/pkg/face"

func (b *Branch) AgeColorFunction() face.SprintFunction {
	return b.AgeColor
}
