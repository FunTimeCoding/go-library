package basic

func New(
	host string,
	user string,
	token string,
	verbose bool,
) *Client {
	// https://developer.atlassian.com/cloud/confluence/rest/v2
	return &Client{
		host:    host,
		user:    user,
		token:   token,
		verbose: verbose,
	}
}
