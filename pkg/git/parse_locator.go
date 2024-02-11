package git

import (
	"github.com/whilp/git-urls"
	"net/url"
)

func ParseLocator(s string) *url.URL {
	webLocator, webLocatorFail := url.Parse(s)

	if webLocatorFail == nil {
		return webLocator
	}

	gitLocator, gitLocatorFail := giturls.Parse(s)

	if gitLocatorFail == nil {
		return gitLocator
	}

	return nil
}
