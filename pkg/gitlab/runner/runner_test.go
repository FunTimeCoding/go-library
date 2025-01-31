package runner

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	gitlab "gitlab.com/gitlab-org/api/client-go"
	"testing"
)

func TestRunner(t *testing.T) {
	r := New(&gitlab.Runner{})
	r.Raw = nil
	assert.Any(t, &Runner{}, r)
}
