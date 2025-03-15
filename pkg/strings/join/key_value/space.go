package key_value

import "fmt"

func Space(
	k string,
	v string,
) string {
	return fmt.Sprintf("%s %s", k, v)
}
