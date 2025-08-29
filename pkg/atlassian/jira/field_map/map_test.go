package field_map

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestMap(t *testing.T) {
	assert.NotNil(t, New([]jira.Field{}))
}
