package content

import (
	"fmt"
	kaos "github.com/essentialkaos/go-confluence/v6"
)

func parseTinyLink(
	v *kaos.Content,
	host string,
) string {
	var result string

	if v.Links != nil {
		result = fmt.Sprintf("https://%s/wiki%s", host, v.Links.TinyUI)
	}

	return result
}
