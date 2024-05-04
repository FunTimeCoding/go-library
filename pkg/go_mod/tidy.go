package go_mod

import "github.com/funtimecoding/go-library/pkg/system"

func Tidy() {
	system.Run("go", "mod", "tidy")
}
