package confluence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) AddComment(
	pageIdentifier string,
	body string,
) {
	payload := commentPayload{
		Type: "comment",
		Container: commentContainer{
			Identifier: pageIdentifier,
			Type:       "page",
		},
		Body: commentBody{
			Storage: commentStorage{
				Value:          fmt.Sprintf("<p>%s</p>", body),
				Representation: "storage",
			},
		},
	}
	c.basic.PostOldPath("/content", notation.Encode(payload, false))
}
