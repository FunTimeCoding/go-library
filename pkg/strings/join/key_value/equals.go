package key_value

import "fmt"

func Equals(
	k string,
	v string,
) string {
	return fmt.Sprintf("%s=%s", k, v)
}
