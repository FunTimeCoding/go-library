package locator

import "fmt"

func (l *Locator) Path(
	f string,
	a ...any,
) *Locator {
	if len(a) > 0 {
		f = fmt.Sprintf(f, a...)
	}

	l.path = f

	return l
}
