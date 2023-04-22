package assert

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	Time(t, now, now)
}
