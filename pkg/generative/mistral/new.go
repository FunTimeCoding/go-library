package mistral

import "github.com/gage-technologies/mistral-go"

func New(token string) *Client {
	return &Client{client: mistral.NewMistralClientDefault(token)}
}
