package key_value

import "fmt"

func Dot(
	k string,
	v string,
) string {
	return fmt.Sprintf("%s.%s", k, v)
}
