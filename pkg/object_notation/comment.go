package object_notation

import "fmt"

func Comment(
	comment string,
	object any,
) string {
	return fmt.Sprintf("%s:\n%s", comment, Format(object))
}
