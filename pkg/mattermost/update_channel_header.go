package mattermost

import (
	"github.com/mattermost/mattermost-server/v6/model"
	"strings"
)

func (c *Client) UpdateChannelHeader(
	h *model.Channel,
	prefix string,
	text string,
	separator string,
) *model.Channel {
	var header []string
	found := false

	for _, s := range strings.Split(h.Header, separator) {
		if strings.HasPrefix(strings.TrimSpace(s), prefix) {
			header = append(header, " "+text+" ")
			found = true
		} else {
			header = append(header, s)
		}
	}

	if !found {
		header[len(header)-1] = header[len(header)-1] + " "
		header = append(header, " "+text)
	}

	h.Header = strings.Join(header, separator)
	result, r, e := c.client.UpdateChannel(h)
	panicOnError(e, r)

	return result
}
