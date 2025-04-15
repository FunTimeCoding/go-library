package basic_client

func New(
	locator string,
	user string,
	password string,
) *Client {
	return &Client{user: user, password: password, locator: locator}
}
