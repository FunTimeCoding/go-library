package locator

import (
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/url"
)

func Stub() *Locator {
	return &Locator{
		scheme:        constant.Secure,
		value:         url.Values{},
		fragmentValue: url.Values{},
	}
}
