package helper

import "github.com/funtimecoding/go-library/pkg/web/locator"

func Base(host string) string {
	return locator.New(host).String()
}
