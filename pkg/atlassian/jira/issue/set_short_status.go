package issue

import "github.com/funtimecoding/go-library/pkg/face"

func (i *Issue) SetShortStatus(f face.StringAlias) {
	i.shortStatus = f
}
