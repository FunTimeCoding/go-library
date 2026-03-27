package naming

import (
	"strings"
	"unicode"
)

func replaceSegment(name, old, replacement string) string {
	spans := segmentSpans(name)
	var target *segmentSpan

	for i := range spans {
		if spans[i].lower == old {
			target = &spans[i]

			break
		}
	}

	if target == nil {
		return name
	}

	firstUpper := unicode.IsUpper(rune(name[target.start]))
	words := strings.Split(replacement, "_")
	underscore := strings.Contains(name, "_")
	var b strings.Builder

	for i, w := range words {
		if i > 0 && underscore {
			b.WriteByte('_')
		}

		if i == 0 && firstUpper {
			b.WriteString(capitalize(w))
		} else if i > 0 && !underscore {
			b.WriteString(capitalize(w))
		} else {
			b.WriteString(w)
		}
	}

	return name[:target.start] + b.String() + name[target.end:]
}

func capitalize(s string) string {
	if s == "" {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])

	return string(r)
}
