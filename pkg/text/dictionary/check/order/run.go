package order

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/text/dictionary"
	"github.com/funtimecoding/go-library/pkg/text/dictionary/constant"
	"sort"
)

func Run() {
	f := constant.File
	c := dictionary.Read(f)

	for i := range c {
		sort.Strings(c[i].Words)
	}

	dictionary.Write(f, c)
	fmt.Printf("Sorted %d categories in %s\n", len(c), f)
}
