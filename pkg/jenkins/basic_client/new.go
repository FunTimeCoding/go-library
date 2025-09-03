package basic_client

func New(
	locator string,
	user string,
	password string,
) *Client {
	return &Client{
		locator:  locator,
		user:     user,
		password: password,
	}
}
