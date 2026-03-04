package assert

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	Duration(t, 5*time.Second, 5*time.Second)
}
