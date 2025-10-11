package basic

func New(
	host string,
	user string,
	token string,
) *Client {
	return &Client{user: user, token: token, host: host}
}
