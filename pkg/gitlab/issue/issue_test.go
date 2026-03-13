package issue

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestIssue(t *testing.T) {
	i := New(&gitlab.Issue{})
	i.Raw = nil
	assert.Any(t, &Issue{}, i)
}
