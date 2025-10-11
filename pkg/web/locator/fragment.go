package locator

import "fmt"

func (l *Locator) Fragment(
	f string,
	a ...any,
) *Locator {
	if len(a) > 0 {
		f = fmt.Sprintf(f, a...)
	}

	l.fragment = f

	return l
}
