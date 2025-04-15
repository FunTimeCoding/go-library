package loader

import "github.com/funtimecoding/go-library/pkg/notation"

func (l *Loader) ToMap() map[string]map[string]string {
	var result = map[string]map[string]string{}

	for k, v := range l.contents {
		var e = map[string]string{}
		notation.DecodeStrict(v, &e, false)
		result[k] = e
	}

	return result
}
