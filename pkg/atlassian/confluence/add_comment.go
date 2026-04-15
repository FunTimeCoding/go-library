package confluence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
)

type commentPayload struct {
	Type      string           `json:"type"`
	Container commentContainer `json:"container"`
	Body      commentBody      `json:"body"`
}

type commentContainer struct {
	Identifier string `json:"id"`
	Type       string `json:"type"`
}

type commentBody struct {
	Storage commentStorage `json:"storage"`
}

type commentStorage struct {
	Value          string `json:"value"`
	Representation string `json:"representation"`
}

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
