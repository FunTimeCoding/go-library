package time

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/constant"
	"testing"
)

func TestFormat(t *testing.T) {
	assert.String(
		t,
		"1970-01-01 00:00",
		Format(constant.StartOfTime),
	)
}
