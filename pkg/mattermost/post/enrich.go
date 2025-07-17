package post

import "github.com/mattermost/mattermost-server/v6/model"

func (p *Post) Enrich(u *model.User) {
	p.User = u
}
