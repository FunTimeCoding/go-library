package pushover

func New(
	user string,
	token string,
) *Client {
	return &Client{user: user, token: token}
}
