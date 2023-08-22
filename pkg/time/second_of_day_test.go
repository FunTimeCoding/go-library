package time

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
	"time"
)

func TestSecondOfDay(t *testing.T) {
	now := time.Now()
	y, m, d := now.Date()
	midnight := time.Date(
		y,
		m,
		d,
		0,
		0,
		0,
		0,
		now.Location(),
	)
	one := time.Date(
		y,
		m,
		d,
		1,
		0,
		0,
		0,
		now.Location(),
	)

	assert.Integer(t, 0, SecondOfDay(midnight))
	assert.Integer(t, 3600, SecondOfDay(one))
}
