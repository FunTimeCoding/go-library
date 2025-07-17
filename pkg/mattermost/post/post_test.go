package post

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/mattermost/mattermost-server/v6/model"
	"testing"
)

func TestPost(t *testing.T) {
	assert.True(t, New(&model.Post{}) != nil)
}
