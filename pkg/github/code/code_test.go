package code

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/google/go-github/v75/github"
	"testing"
)

func TestCode(t *testing.T) {
	r := New(
		&github.CodeResult{
			SHA:  ptr.To(strings.Alfa),
			Name: ptr.To(strings.Bravo),
			Path: ptr.To(strings.Charlie),
		},
	)
	r.Raw = nil
	assert.Any(t, &Code{Hash: "Alfa", Name: "Bravo", Path: "Charlie"}, r)
}
