package basic

import "github.com/funtimecoding/go-library/pkg/web/locator"

type Client struct {
	host    string
	user    string
	token   string
	verbose bool
	base    *locator.Locator
}
