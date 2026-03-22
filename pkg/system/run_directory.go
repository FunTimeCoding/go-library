package system

import "github.com/funtimecoding/go-library/pkg/system/run"

func RunDirectory(
	directory string,
	s ...string,
) string {
	r := run.New()
	r.Directory = directory

	return r.Start(s...)
}
