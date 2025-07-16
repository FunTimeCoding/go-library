package n8n

func New(
	host string,
	token string,
) *Client {
	return &Client{host: host, token: token}
}
