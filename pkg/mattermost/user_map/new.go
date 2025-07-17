package user_map

import "github.com/mattermost/mattermost-server/v6/model"

func New() *Map {
	return &Map{
		byDirectory:  make(map[string]*model.User),
		byIdentifier: make(map[string]*model.User),
	}
}
