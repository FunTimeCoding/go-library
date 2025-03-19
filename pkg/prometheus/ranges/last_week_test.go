package ranges

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
	"time"
)

func TestLastWeek(t *testing.T) {
	assert.Any(t, time.Minute, LastWeek(1).Step)
}
