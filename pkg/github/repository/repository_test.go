package repository

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/google/go-github/v70/github"
	"testing"
	"time"
)

func TestRepository(t *testing.T) {
	r := New(
		&github.Repository{
			Name:      ptr.To(strings.Alfa),
			CreatedAt: &github.Timestamp{},
		},
	)
	r.Raw = nil
	assert.Any(t, &Repository{Name: "Alfa", CreatedAt: time.Time{}}, r)
}
