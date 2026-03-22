package system

import "github.com/funtimecoding/go-library/pkg/system/run"

func Pipe(
	input string,
	s ...string,
) (string, string) {
	return run.New().Pipe(input, s...)
}
