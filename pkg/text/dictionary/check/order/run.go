package order

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/text/dictionary"
	"sort"
)

func Run() {
	f := dictionary.ResolvePath()
	c := dictionary.Read(f)

	for i := range c {
		sort.Strings(c[i].Words)
	}

	dictionary.Write(f, c)
	fmt.Printf("Sorted %d categories in %s\n", len(c), f)
}
