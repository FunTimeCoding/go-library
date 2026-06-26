package page_put

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
)

func NewWithStatus(
	identifier string,
	title string,
	body string,
	version int,
	message string,
	status string,
) *Put {
	return &Put{
		Identifier: identifier,
		Status:     status,
		Title:      title,
		Body: response.Storage{
			Representation: constant.StorageFormat,
			Value:          body,
		},
		Version: Version{
			Number:  version,
			Message: message,
		},
	}
}
