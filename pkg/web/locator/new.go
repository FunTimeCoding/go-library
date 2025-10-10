package locator

import (
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/url"
)

func New() *Locator {
	return &Locator{
		values: url.Values{},
		scheme: constant.SecureScheme,
	}
}
