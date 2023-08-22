package time

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestHumanReadable(t *testing.T) {
	minute := MinuteInSeconds
	hour := HourInSeconds
	day := DayInSeconds
	month := MonthInSeconds

	assert.String(t, "1 second", HumanReadable(1))
	assert.String(t, "59 seconds", HumanReadable(59))
	assert.String(t, "1 minute", HumanReadable(minute))

	assert.String(t, "30 minutes", HumanReadable(minute*30))
	assert.String(t, "59 minutes", HumanReadable(minute*59))
	assert.String(t, "1 hour", HumanReadable(hour))
	assert.String(t, "1 day", HumanReadable(day))
	assert.String(t, "1.5 days", HumanReadable(hour*36))
	assert.String(t, "2 days", HumanReadable(day*2))
	assert.String(t, "1 week", HumanReadable(day*7))
	assert.String(t, "1 month", HumanReadable(month))
}
