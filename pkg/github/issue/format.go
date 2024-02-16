package issue

import "fmt"

func (i *Issue) Format(raw bool) string {
	result := fmt.Sprintf("%s | %s", i.Repository, i.Title)

	if raw {
		result = fmt.Sprintf("%s | %+v", result, i.Raw)
	}

	return result
}
