package sentry

func (c *Client) Favourites() string {
	return c.basic.Get("/users/favorites/")
}
