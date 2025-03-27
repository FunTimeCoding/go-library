package basic_client

func New(
	host string,
	user string,
	token string,
) *Client {
	return &Client{host: host, user: user, token: token}
}
