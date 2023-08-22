package time

import (
	assert2 "github.com/funtimecoding/go-library/pkg/assert"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	assert2.Integer(t, 3600, HourInSeconds)
	assert2.Integer(t, 86400, DayInSeconds)
	assert2.Integer(t, 604800, WeekInSeconds)
	assert2.Integer(t, 2419200, MonthInSeconds)
	assert2.Integer(t, 29030400, YearInSeconds)

	assert2.Integer(t, 28, MonthInDays)

	assert2.String(
		t,
		"2022-01-01 00:00:00",
		time.Date(
			2022,
			time.January,
			1,
			0,
			0,
			0,
			0,
			time.UTC,
		).Format("2006-01-02 15:04:05"),
	)

	now := time.Now()
	past := now.Add(-time.Minute)

	if !past.Before(now) {
		t.Fatalf("Past not before present")
	}

	future := now.Add(time.Minute)

	if !future.After(now) {
		t.Fatalf("Future not after present")
	}
}
