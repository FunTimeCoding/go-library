package key_value

import "fmt"

func Slash(
	k string,
	v string,
) string {
	return fmt.Sprintf("%s/%s", k, v)
}
