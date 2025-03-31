package user

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"
	"testing"
)

func TestUser(t *testing.T) {
	assert.True(t, New(&response.User{}) != nil)
}
