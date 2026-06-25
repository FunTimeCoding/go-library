package technitium

import "net/http"

type Client struct {
	base   string
	token  string
	client *http.Client
}
