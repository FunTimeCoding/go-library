package release

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/google/go-github/v79/github"
	"testing"
	"time"
)

func TestRelease(t *testing.T) {
	r := New(
		&github.RepositoryRelease{
			TagName:   ptr.To(strings.Alfa),
			CreatedAt: &github.Timestamp{},
		},
	)
	r.Raw = nil
	assert.Any(
		t,
		&Release{Name: "Alfa", Create: time.Time{}},
		r,
	)
}
