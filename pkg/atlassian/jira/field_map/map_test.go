package field_map

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestMap(t *testing.T) {
	assert.True(t, New([]jira.Field{}) != nil)
}
