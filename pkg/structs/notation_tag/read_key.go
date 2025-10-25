package notation_tag

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/structs/constant"
	"log"
)

func ReadKey(v []string) string {
	fmt.Printf("Before ReadKey: %v\n", v)
	v = strings.RemoveFromList(v, []string{constant.OmitEmpty})
	fmt.Printf("After ReadKey: %v\n", v)

	if len(v) == 0 {
		panic("empty notation")
	}

	if len(v) > 1 {
		log.Panicf("too many values: %v", v)
	}

	return v[0]
}
