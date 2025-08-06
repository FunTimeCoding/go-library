package post

import "github.com/mattermost/mattermost/server/public/model"

func (p *Post) Enrich(u *model.User) {
	p.User = u
}
