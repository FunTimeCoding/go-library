package issue

import "github.com/funtimecoding/go-library/pkg/face"

func (i *Issue) SetScoreColor(f face.SprintFunction) {
	i.scoreColor = f
}
