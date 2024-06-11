package argument

import (
	"fmt"
	"os"
)

func RequiredPositional(
	number int,
	description string,
	exitCode int,
) string {
	if s := Positional(number); s != "" {
		return s
	}

	fmt.Printf("positional argument empty: %s\n", description)
	os.Exit(exitCode)

	return ""
}
