package thread

import "github.com/funtimecoding/go-library/pkg/chat/mattermost/post"

type Thread struct {
	Root     *post.Post
	Posts    []*post.Post
	Resolved bool
}
