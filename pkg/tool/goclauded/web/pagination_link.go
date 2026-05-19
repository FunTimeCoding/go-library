package web

import "fmt"

func paginationLink(page int) string {
	if page <= 1 {
		return "/history"
	}

	return fmt.Sprintf("/history?page=%d", page)
}
