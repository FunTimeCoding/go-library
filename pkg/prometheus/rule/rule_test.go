package rule

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestRule(t *testing.T) {
	assert.True(t, NewAlert(nil, nil) != nil)
	assert.True(t, NewRecord(nil, nil) != nil)
}
