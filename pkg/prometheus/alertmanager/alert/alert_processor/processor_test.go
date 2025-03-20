package alert_processor

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestProcessor(t *testing.T) {
	assert.True(t, New(nil, nil, nil, nil, nil, nil) != nil)
}
