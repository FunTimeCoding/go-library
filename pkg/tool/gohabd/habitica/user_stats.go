package habitica

import "net/http"

func (c *Client) UserStats() Stats {
	var user struct {
		Stats Stats `json:"stats"`
	}
	r := c.do(http.MethodGet, "/user", nil)
	c.decode(r, &user)

	return user.Stats
}
