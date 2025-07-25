package tag

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"gitlab.com/gitlab-org/api/client-go"
	"testing"
)

func TestLatest(t *testing.T) {
	assert.Any(
		t,
		&gitlab.Tag{Name: "1.0.0"},
		Latest([]*gitlab.Tag{{Name: "1.0.0"}}),
	)
}
