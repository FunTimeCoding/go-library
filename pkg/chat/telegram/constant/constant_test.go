package constant

import (
	"testing"

	"github.com/funtimecoding/go-library/pkg/assert"
)

func TestConstant(t *testing.T) {
	assert.String(t, "TELEGRAM_TOKEN", TokenEnvironment)
}
