package post

import (
	"time"

	"github.com/mattermost/mattermost/server/public/model"
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
