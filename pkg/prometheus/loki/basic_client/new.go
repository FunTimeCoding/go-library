package basic_client

func New(
	host string,
	user string,
	password string,
) *Client {
	return &Client{host: host, user: user, password: password}
}
