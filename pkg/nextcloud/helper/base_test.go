package helper

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"testing"
)

func TestBase(t *testing.T) {
	assert.String(t, "https://example.org", Base(constant.Example))
}
