package argument

import (
	"fmt"
	"os"
)

func RequiredPositional(
	number int,
	description string,
) string {
	if s := Positional(number); s != "" {
		return s
	}

	fmt.Printf("required positional empty: %s\n", description)
	os.Exit(1)

	return ""
}
