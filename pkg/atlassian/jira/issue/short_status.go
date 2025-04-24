package issue

import "github.com/funtimecoding/go-library/pkg/face"

func (i *Issue) ShortStatus() face.StringAlias {
	return i.shortStatus
}
