package notation_tag

import (
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/structs/constant"
	"log"
)

func ReadKey(v []string) string {
	v = strings.RemoveFromList(v, []string{constant.OmitEmpty})

	if len(v) == 0 {
		panic("empty notation")
	}

	if len(v) > 1 {
		log.Panicf("too many values: %v", v)
	}

	return v[0]
}
