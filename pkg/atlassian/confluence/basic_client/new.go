package basic_client

func New(
	host string,
	user string,
	token string,
) *Client {
	// https://developer.atlassian.com/cloud/confluence/rest/v2
	return &Client{host: host, user: user, token: token}
}
