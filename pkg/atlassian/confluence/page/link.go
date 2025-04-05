package page

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"
)

func link(
	p *response.Page,
	host string,
	tiny bool,
) string {
	if tiny {
		if p.Links.WebUI != "" {
			return fmt.Sprintf(
				"https://%s/wiki%s",
				host,
				p.Links.TinyUI,
			)
		}
	} else {
		if p.Links.WebUI != "" {
			return fmt.Sprintf(
				"https://%s/wiki%s",
				host,
				p.Links.WebUI,
			)
		}
	}

	return ""
}
