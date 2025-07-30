package post

import (
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

type Post struct {
	Identifier     string
	UserIdentifier string
	Message        string
	Create         time.Time

	User *model.User

	Raw *model.Post
}
