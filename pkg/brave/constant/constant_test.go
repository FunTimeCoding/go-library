package constant

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(
		t,
		"/Applications/Brave Browser.app/Contents/MacOS/Brave Browser",
		BraveMacPath,
	)
}
