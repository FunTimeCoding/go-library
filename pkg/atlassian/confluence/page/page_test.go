package page

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestPage(t *testing.T) {
	assert.NotNil(t, New(response.NewPage(), upper.Alfa))
}
