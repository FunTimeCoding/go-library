package argument

import "github.com/spf13/pflag"

// Positional returns empty string if out of bounds
func Positional(number int) string {
	return pflag.Arg(number)
}
