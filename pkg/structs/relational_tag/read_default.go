package relational_tag

import (
	"github.com/funtimecoding/go-library/pkg/strings/split/key_value"
	"github.com/funtimecoding/go-library/pkg/structs/constant"
	"strings"
)

func ReadDefault(v []string) string {
	for _, e := range v {
		if strings.HasPrefix(e, constant.DefaultPrefix) {
			_, a := key_value.Colon(e)

			return a
		}
	}

	return ""
}
