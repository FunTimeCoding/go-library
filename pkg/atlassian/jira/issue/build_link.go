package issue

import "fmt"

func BuildLink(
	locator string,
	key string,
) string {
	return fmt.Sprintf("%s/browse/%s", locator, key)
}
