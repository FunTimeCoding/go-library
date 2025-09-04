package basic_client

func New(
	locator string,
	user string,
	token string,
) *Client {
	return &Client{user: user, token: token, locator: locator}
}
