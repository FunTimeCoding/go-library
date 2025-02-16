package helper

import (
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestHelper(t *testing.T) {
	ValidateContains([]string{strings.Alfa, strings.Bravo}, "Alfa")
}
