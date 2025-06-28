package job

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"gitlab.com/gitlab-org/api/client-go"
	"testing"
)

func TestDeduplicate(t *testing.T) {
	assert.Count(t, 0, Deduplicate([]*gitlab.Job{}))
}
