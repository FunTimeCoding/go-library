package post

import (
	"testing"

	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/mattermost/mattermost/server/public/model"
)

func TestPost(t *testing.T) {
	assert.True(t, New(&model.Post{}) != nil)
}
