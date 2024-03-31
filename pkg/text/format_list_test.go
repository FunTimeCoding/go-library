package text

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestFormatList(t *testing.T) {
	assert.String(
		t,
		FormatList("title", []string{"element"}),
		"title:\nelement",
	)
}
