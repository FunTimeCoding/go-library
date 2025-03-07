package branch

import "github.com/funtimecoding/go-library/pkg/face"

func (b *Branch) SetAgeColorFunction(f face.SprintFunction) {
	b.AgeColor = f
}
