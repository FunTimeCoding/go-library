package option

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestOption(t *testing.T) {
	assert.NotNil(
		t,
		New(
			upper.Alfa,
			upper.Bravo,
			[]string{},
			[]string{constant.Done},
			nil,
		),
	)
}
