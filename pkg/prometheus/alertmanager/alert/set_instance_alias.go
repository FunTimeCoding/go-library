package alert

import "github.com/funtimecoding/go-library/pkg/face"

func (a *Alert) SetInstanceAlias(f face.StringAlias) {
	a.instance = f
}
