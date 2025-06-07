package gmail

import (
	"github.com/funtimecoding/go-library/pkg/gmail/constant"
	"github.com/funtimecoding/go-library/pkg/system"
	"golang.org/x/oauth2"
	"strings"
)

func (c *Client) loadTokens() map[string]*oauth2.Token {
	result := make(map[string]*oauth2.Token)

	for _, n := range system.DirectoryContent(c.directory) {
		if strings.HasSuffix(n, constant.TokenSuffix) {
			result[n] = c.loadToken(system.Join(c.directory, n))
		}
	}

	return result
}
