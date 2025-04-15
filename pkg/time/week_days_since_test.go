package time

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
	"time"
)

func TestWorkDaysSince(t *testing.T) {
	assert.Integer(
		t,
		3,
		WeekDaysSince(
			NewDay(2024, time.October, 25),
			NewDay(2024, time.October, 30),
		),
	)
}
