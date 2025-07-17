package post

import "github.com/mattermost/mattermost-server/v6/model"

func FromList(l *model.PostList) []*model.Post {
	var result []*model.Post

	for _, v := range l.Posts {
		result = append(result, v)
	}

	Sort(result)

	return result
}
