package web

import "fmt"

func messagePaginationLink(page int) string {
	if page <= 1 {
		return "/messages"
	}

	return fmt.Sprintf("/messages?page=%d", page)
}
