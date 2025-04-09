package page_post

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"

type Post struct {
	SpaceIdentifier  string           `json:"spaceId"`
	Status           string           `json:"status"`
	Title            string           `json:"title"`
	ParentIdentifier string           `json:"parentId"`
	Body             response.Storage `json:"body"`
}
