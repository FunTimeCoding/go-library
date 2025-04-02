package loader

import "github.com/funtimecoding/go-library/pkg/notation"

func (l *Loader) ToMap() map[string]map[string]string {
	var result = map[string]map[string]string{}

	for k, v := range l.contents {
		var element = map[string]string{}
		notation.DecodeStrict(v, &element, false)
		result[k] = element
	}

	return result
}
