package mock_notifier

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestNotifier(t *testing.T) {
	assert.NotNil(t, New(""))
}
