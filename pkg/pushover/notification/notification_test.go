package notification

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/notation"
	"testing"
)

func TestNotification(t *testing.T) {
	n := New("a", "b", "c")
	n.SetPriority(1)
	assert.String(
		t,
		`{"user":"a","token":"b","message":"c","priority":1}`,
		notation.Encode(n, false),
	)
}
