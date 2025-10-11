package page_post

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
)

func New(
	spaceIdentifier string,
	parentIdentifier string,
	title string,
	body string,
) *Post {
	return &Post{
		SpaceIdentifier:  spaceIdentifier,
		ParentIdentifier: parentIdentifier,
		Title:            title,
		Status:           "current",
		Body: response.Storage{
			Representation: constant.StorageFormat,
			Value:          body,
		},
	}
}
