package sentry

func (c *Client) Favourites() string {
	return c.basicClient.Request("/users/favorites/")
}
