package go_mod

import "github.com/funtimecoding/go-library/pkg/system"

func Update(name string) {
	system.Run("go", "get", name)
}
