package basic

func New(
	host string,
	port int,
	user string,
	password string,
) *Client {
	return &Client{
		host:     host,
		port:     port,
		user:     user,
		password: password,
	}
}
