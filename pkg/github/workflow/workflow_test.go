package workflow

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/google/go-github/v82/github"
	"testing"
	"time"
)

func TestWorkflow(t *testing.T) {
	r := New(
		&github.Workflow{
			Name:      ptr.To(strings.Alfa),
			CreatedAt: &github.Timestamp{},
		},
	)
	r.Raw = nil
	assert.Any(t, &Workflow{Name: "Alfa", CreatedAt: time.Time{}}, r)
}
