package loader

import (
	"github.com/funtimecoding/go-library/pkg/maps"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (l *Loader) ToText() string {
	var result []string

	for k, v := range l.ToMap() {
		result = append(result, IntegerFromString(k))

		for _, kk := range maps.StringKeys(v) {
			result = append(result, kk+": "+v[kk])
		}
	}

	return join.NewLine(result)
}
