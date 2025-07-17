package user_map

import "github.com/mattermost/mattermost-server/v6/model"

type Map struct {
	byDirectory  map[string]*model.User
	byIdentifier map[string]*model.User
}
