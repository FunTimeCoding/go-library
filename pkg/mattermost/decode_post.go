package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mattermost/mattermost/server/public/model"
)

func DecodePost(v *model.WebSocketEvent) *model.Post {
	raw := v.GetData()["post"].(string)
	var p *model.Post
	notation.DecodeStrict(raw, &p, false)

	return p
}
