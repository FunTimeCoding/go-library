package errors

import (
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"testing"
)

func TestPost(t *testing.T) {
	Post(
		locator.New(constant.Localhost).String(),
		"something went wrong",
	)
}
