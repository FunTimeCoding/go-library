package ranges

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
	"time"
)

func TestLastDay(t *testing.T) {
	assert.Any(t, time.Minute, LastDay(1).Step)
}
