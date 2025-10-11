package content

import (
	kaos "github.com/essentialkaos/go-confluence/v6"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func parseTinyLink(
	v *kaos.Content,
	host string,
) string {
	var result string

	if v.Links != nil {
		result = locator.New(
			host,
		).Base(constant.Wiki).Path(v.Links.TinyUI).String()
	}

	return result
}
