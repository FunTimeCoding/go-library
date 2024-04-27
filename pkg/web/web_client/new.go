package web_client

import "github.com/funtimecoding/go-library/pkg/face"

func New(c face.Clock) *Client {
	return &Client{clock: c}
}
