package post

import "github.com/mattermost/mattermost/server/public/model"

func FromList(
	l *model.PostList,
	oldestFirst bool,
) []*model.Post {
	var result []*model.Post

	for _, v := range l.Posts {
		result = append(result, v)
	}

	if oldestFirst {
		Sort(result)
	}

	return result
}
