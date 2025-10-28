package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/brave"
	"github.com/funtimecoding/go-library/pkg/system"
)

func Bookmark() {
	for _, p := range brave.Profiles() {
		fmt.Printf("Profile: %+v\n", p)
		// TODO: Read bookmarks

		b := system.ReadFile(p.Path, "Bookmarks")
		fmt.Printf("%s\n", b)
	}
}
