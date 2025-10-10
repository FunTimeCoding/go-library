package locator

import "github.com/funtimecoding/go-library/pkg/web/constant"

func (l *Locator) Insecure() *Locator {
	l.scheme = constant.InsecureScheme

	return l
}
