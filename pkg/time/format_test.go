package time

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	assert.String(
		t,
		"2022-01-01 00:00",
		Format(
			time.Date(
				2022,
				time.January,
				1,
				0,
				0,
				0,
				0,
				time.UTC,
			),
		),
	)
}
