package locator

import "github.com/funtimecoding/go-library/pkg/integers"

func (l *Locator) SetInteger(
	k string,
	v int,
) *Locator {
	l.value.Set(k, integers.ToString(v))

	return l
}
