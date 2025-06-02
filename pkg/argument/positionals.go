package argument

import "github.com/spf13/pflag"

func Positionals() []string {
	return pflag.Args()
}
