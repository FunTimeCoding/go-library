package chromium

import "context"

func (c *Client) FillNode(
	x context.Context,
	backendNodeID int64,
	value string,
	direct bool,
) error {
	if direct {
		return c.fillNodeDirect(x, backendNodeID, value)
	}

	return c.fillNodeInsertText(x, backendNodeID, value)
}
