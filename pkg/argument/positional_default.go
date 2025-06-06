package argument

import "github.com/spf13/pflag"

func PositionalFallback(
	number int,
	fallback string,
) string {
	value := pflag.Arg(number)

	if value == "" {
		return fallback
	}

	return value
}
