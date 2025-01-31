package commit

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	gitlab "gitlab.com/gitlab-org/api/client-go"
	"testing"
)

func TestCommit(t *testing.T) {
	c := New(&gitlab.Commit{})
	c.Raw = nil
	assert.Any(t, &Commit{}, c)
}
