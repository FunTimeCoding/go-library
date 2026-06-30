package run

import (
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestValidateArgument(t *testing.T) {
	ValidateArgument(upper.Alfa)
}
