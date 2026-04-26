package web

import "fmt"

func paginationLink(
	offset int,
	startValue, endValue string,
) string {
	link := fmt.Sprintf("/?offset=%d", offset)

	if startValue != "" {
		link = fmt.Sprintf("%s&start=%s", link, startValue)
	}

	if endValue != "" {
		link = fmt.Sprintf("%s&end=%s", link, endValue)
	}

	return link
}
