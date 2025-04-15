package query

import (
	"fmt"
	"strings"
)

func Quote(v []string) []string {
	for i, e := range v {
		if strings.ContainsRune(e, ' ') {
			v[i] = fmt.Sprintf("'%s'", e)
		}
	}

	return v
}
