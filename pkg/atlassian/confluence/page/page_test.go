package page

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestPage(t *testing.T) {
	assert.NotNil(t, New(&response.Page{}, strings.Alfa))
}
