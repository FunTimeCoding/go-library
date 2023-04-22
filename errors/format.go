package errors

import "fmt"

func Format(
	text string,
	detail any,
) error {
	return fmt.Errorf("%s: %#v", text, detail)
}
