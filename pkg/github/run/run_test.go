package run

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/google/go-github/v81/github"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	r := New(
		&github.WorkflowRun{
			Name:       ptr.To(strings.Alfa),
			CreatedAt:  &github.Timestamp{},
			Repository: &github.Repository{},
		},
	)
	r.Repository = nil
	r.Raw = nil
	assert.Any(
		t,
		&Run{
			MonitorIdentifier: "ghjob-0",
			Name:              "Alfa",
			Create:            time.Time{},
		},
		r,
	)
}
