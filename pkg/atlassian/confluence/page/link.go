package page

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func link(
	p *response.Page,
	host string,
	tiny bool,
) string {
	if tiny {
		if p.Links.WebUI != "" {
			return locator.New(
				host,
			).Base(constant.Wiki).Path(p.Links.TinyUI).String()
		}
	} else {
		if p.Links.WebUI != "" {
			return locator.New(
				host,
			).Base(constant.Wiki).Path(p.Links.WebUI).String()
		}
	}

	return ""
}
