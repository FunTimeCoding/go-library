package user_map

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/user"
	"testing"
)

func TestMap(t *testing.T) {
	assert.True(t, New([]*user.User{}) != nil)
}
