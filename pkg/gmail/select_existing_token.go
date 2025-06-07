package gmail

import "golang.org/x/oauth2"

func (c *Client) selectExistingToken() *oauth2.Token {
	var result *oauth2.Token
	tokens := c.loadTokens()
	count := len(tokens)

	if count > 0 {
		if count == 1 {
			for _, t := range tokens {
				result = t
			}
		} else {
			result = tokenSelector(tokens)
		}
	}

	return result
}
