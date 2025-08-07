package post

import (
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func New(p *model.Post) *Post {
	return &Post{
		Identifier:     p.Id,
		UserIdentifier: p.UserId,
		Message:        p.Message,

		Create: time.UnixMilli(p.CreateAt),

		Raw: p,
	}
}
