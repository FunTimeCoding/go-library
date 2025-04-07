package openapi

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/go-openapi/strfmt"
	"testing"
	"time"
)

func TestConvertTime(t *testing.T) {
	i := strfmt.NewDateTime()
	errors.PanicOnError(i.Scan(constant.StartOfTime))
	assert.String(
		t,
		"1970-01-01T00:00:00Z",
		ConvertTime(&i).Format(time.RFC3339),
	)
}
