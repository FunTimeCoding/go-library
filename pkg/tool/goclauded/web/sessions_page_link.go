package web

import "fmt"

func sessionsPageLink(
	base string,
	page int,
) string {
	if page <= 1 {
		return base
	}

	separator := "?"

	if len(base) > 0 && base[len(base)-1] != '/' {
		for _, c := range base {
			if c == '?' {
				separator = "&"

				break
			}
		}
	}

	return fmt.Sprintf("%s%spage=%d", base, separator, page)
}
