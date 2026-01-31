package markdown

import "fmt"

func Link(
	title string,
	locator string,
) string {
	return fmt.Sprintf("[%s](%s)", title, locator)
}
