package web

import "fmt"

func paginationLink(offset int, startValue, endValue string) string {
	link := fmt.Sprintf("/?offset=%d", offset)

	if startValue != "" {
		link += fmt.Sprintf("&start=%s", startValue)
	}

	if endValue != "" {
		link += fmt.Sprintf("&end=%s", endValue)
	}

	return link
}
