package post

import "github.com/mattermost/mattermost/server/public/model"

func NewSlice(v []*model.Post) []*Post {
	var result []*Post

	for _, p := range v {
		result = append(result, New(p))
	}

	return result
}
