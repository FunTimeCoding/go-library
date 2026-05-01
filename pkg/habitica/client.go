package habitica

import "net/http"

type Client struct {
	base           string
	userIdentifier string
	token          string
	client         *http.Client
}
