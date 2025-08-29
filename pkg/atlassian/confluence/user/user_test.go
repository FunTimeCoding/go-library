package user

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"
	"testing"
)

func TestUser(t *testing.T) {
	assert.NotNil(t, New(&response.User{}))
}
