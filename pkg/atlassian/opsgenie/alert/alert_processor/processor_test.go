package alert_processor

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestProcessor(t *testing.T) {
	assert.NotNil(t, New(nil))
}
