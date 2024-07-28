package runtime

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func Fixture() {
}

func TestFunctionName(t *testing.T) {
	assert.String(
		t,
		"github.com/funtimecoding/go-library/pkg/runtime.Fixture",
		FunctionName(Fixture),
	)
}
