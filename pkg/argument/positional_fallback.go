package argument

import "github.com/spf13/pflag"

func PositionalFallback(
	number int,
	fallback string,
) string {
	if s := pflag.Arg(number); s == "" {
		return s
	}

	return fallback
}
