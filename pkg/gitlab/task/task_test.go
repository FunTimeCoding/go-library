package task

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	gitlab "gitlab.com/gitlab-org/api/client-go"
	"testing"
)

func TestTask(t *testing.T) {
	a := New(&gitlab.Todo{})
	a.Raw = nil
	assert.Any(t, &Task{}, a)
}
