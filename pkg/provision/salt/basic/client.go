package basic

import "net/http"

type Client struct {
	base     string
	token    string
	user     string
	password string
	eauth    string
	client   *http.Client
}
