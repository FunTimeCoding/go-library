package workflow

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/google/go-github/v85/github"
	"testing"
	"time"
)

func TestWorkflow(t *testing.T) {
	r := New(
		&github.Workflow{
			Name:      new(strings.Alfa),
			CreatedAt: &github.Timestamp{},
		},
	)
	r.Raw = nil
	assert.Any(t, &Workflow{Name: "Alfa", CreatedAt: time.Time{}}, r)
}
