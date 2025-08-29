package option

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestOption(t *testing.T) {
	assert.NotNil(
		t,
		New(
			strings.Alfa,
			strings.Bravo,
			[]string{},
			[]string{constant.Done},
			nil,
		),
	)
}
