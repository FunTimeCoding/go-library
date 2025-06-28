package runner

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"gitlab.com/gitlab-org/api/client-go"
	"testing"
)

func TestRunner(t *testing.T) {
	r := New(&gitlab.Runner{})
	r.RawList = nil
	assert.Any(t, &Runner{}, r)
}
