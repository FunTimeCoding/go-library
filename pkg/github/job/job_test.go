package job

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/google/go-github/v80/github"
	"testing"
	"time"
)

func TestJob(t *testing.T) {
	r := New(
		&github.WorkflowJob{
			Name:      ptr.To(strings.Alfa),
			CreatedAt: &github.Timestamp{},
		},
	)
	r.Raw = nil
	assert.Any(t, &Job{Name: "Alfa", CreatedAt: time.Time{}}, r)
}
