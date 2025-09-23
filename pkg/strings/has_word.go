package strings

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"regexp"
)

func HasWord(
	text string,
	word string,
) bool {
	// matches word with word boundaries
	// prevents partial matches like "test" matching "testing"
	result, e := regexp.MatchString(
		`\b`+regexp.QuoteMeta(word)+`\b`,
		text,
	)
	errors.PanicOnError(e)

	return result
}
