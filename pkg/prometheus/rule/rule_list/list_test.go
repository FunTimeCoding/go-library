package rule_list

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestList(t *testing.T) {
	assert.True(t, New() != nil)
}
