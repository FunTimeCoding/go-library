package kubernetes

import (
	"github.com/funtimecoding/go-library/pkg/strings/lower"
	"testing"
)

func TestValidateName(t *testing.T) {
	ValidateName(lower.Alfa)
}
