package locator

import "fmt"

func (l *Locator) Base(
	f string,
	a ...any,
) *Locator {
	if len(a) > 0 {
		f = fmt.Sprintf(f, a...)
	}

	l.basePath = f

	return l
}
