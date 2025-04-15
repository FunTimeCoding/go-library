package time

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
	"time"
)

func TestPublicHoliday(t *testing.T) {
	assert.True(t, PublicHoliday(NewDay(2024, time.November, 1)))
	assert.False(t, PublicHoliday(NewDay(2024, time.November, 2)))
}
