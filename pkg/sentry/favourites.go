package sentry

func (c *Client) Favourites() string {
	return c.basicClient.Get("/users/favorites/")
}
