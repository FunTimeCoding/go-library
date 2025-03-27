package content

import kaos "github.com/essentialkaos/go-confluence/v6"

func parseSpace(v *kaos.Content) string {
	if v.Space != nil {
		return v.Space.Name
	}

	return ""
}
