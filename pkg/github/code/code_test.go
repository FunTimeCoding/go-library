package code

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/google/go-github/v86/github"
	"testing"
)

func TestCode(t *testing.T) {
	r := New(
		&github.CodeResult{
			SHA:  new(strings.Alfa),
			Name: new(strings.Bravo),
			Path: new(strings.Charlie),
		},
	)
	r.Raw = nil
	assert.Any(t, &Code{Hash: "Alfa", Name: "Bravo", Path: "Charlie"}, r)
}
