package key_value

import "fmt"

func Colon(
	k string,
	v string,
) string {
	return fmt.Sprintf("%s:%s", k, v)
}
