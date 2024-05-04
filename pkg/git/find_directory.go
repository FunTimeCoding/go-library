package git

import "github.com/funtimecoding/go-library/pkg/system"

func FindDirectory() string {
	return system.FindDirectoryUp(system.WorkingDirectory(), Directory)
}
