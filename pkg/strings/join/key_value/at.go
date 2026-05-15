package key_value

import "fmt"

func At(
	k string,
	v string,
) string {
	return fmt.Sprintf("%s@%s", k, v)
}
