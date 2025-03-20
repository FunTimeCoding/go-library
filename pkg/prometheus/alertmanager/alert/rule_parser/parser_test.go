package rule_parser

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestParser(t *testing.T) {
	assert.True(t, New(nil) != nil)
}
