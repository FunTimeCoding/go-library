package openapi

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/go-openapi/strfmt"
	"testing"
	"time"
)

func TestConvertTime(t *testing.T) {
	i := strfmt.NewDateTime()
	errors.PanicOnError(
		i.Scan(
			time.Date(
				2021,
				1,
				1,
				0,
				0,
				0,
				0,
				time.UTC,
			),
		),
	)
	assert.String(
		t,
		"2021-01-01T00:00:00Z",
		ConvertTime(&i).Format(time.RFC3339),
	)
}
