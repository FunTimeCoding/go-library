package argument

import "github.com/funtimecoding/go-library/pkg/system"

func RequiredPositionalExit(
	number int,
	description string,
	exit int,
) string {
	if s := Positional(number); s != "" {
		return s
	}

	system.Exitf(
		exit,
		"positional argument empty: %s\n",
		description,
	)

	return ""
}
