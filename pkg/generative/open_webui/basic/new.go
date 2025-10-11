package basic

func New(
	host string,
	token string,
) *Client {
	return &Client{host: host, token: token}
}
