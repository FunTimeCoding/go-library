package argument

import "github.com/funtimecoding/go-library/pkg/system"

func RequiredPositional(
	number int,
	description string,
	exitCode int,
) string {
	if s := Positional(number); s != "" {
		return s
	}

	system.Exitf(
		exitCode,
		"positional argument empty: %s\n",
		description,
	)

	return ""
}
