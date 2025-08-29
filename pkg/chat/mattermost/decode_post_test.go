package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/mattermost/mattermost/server/public/model"
	"testing"
)

func TestDecodePost(t *testing.T) {
	e := &model.WebSocketEvent{}
	e = e.SetData(map[string]any{"post": `{"id":"alfa"}`})
	assert.Any(t, &model.Post{Id: "alfa"}, DecodePost(e))
}
