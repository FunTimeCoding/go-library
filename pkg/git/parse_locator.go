package git

import (
	"github.com/whilp/git-urls"
	"net/url"
)

func ParseLocator(s string) *url.URL {
	if l, e := url.Parse(s); e == nil {
		return l
	}

	if l, e := giturls.Parse(s); e == nil {
		return l
	}

	return nil
}
