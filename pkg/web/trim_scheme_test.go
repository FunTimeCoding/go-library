package web

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"testing"
)

func TestTrimScheme(t *testing.T) {
	assert.String(
		t,
		"localhost",
		TrimScheme(locator.New(constant.Localhost).String()),
	)
	assert.String(
		t,
		"localhost",
		TrimScheme(locator.New(constant.Localhost).Insecure().String()),
	)
}
