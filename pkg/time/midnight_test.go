package time

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/constant"
	"testing"
	"time"
)

func TestMidnight(t *testing.T) {
	assert.Any(
		t, time.Date(
			1970,
			1,
			1,
			0,
			0,
			0,
			0,
			time.Local,
		),
		Midnight(constant.StartOfTime),
	)
}
