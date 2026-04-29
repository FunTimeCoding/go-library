package habitica

import "net/http"

type Client struct {
	baseURL string
	userID  string
	token   string
	http    *http.Client
}
