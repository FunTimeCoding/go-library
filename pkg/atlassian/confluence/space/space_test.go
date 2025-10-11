package space

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"testing"
)

func TestSpace(t *testing.T) {
	assert.NotNil(t, New(&response.Space{}))
}
