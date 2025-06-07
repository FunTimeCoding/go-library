package gmail

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gmail/constant"
	"github.com/funtimecoding/go-library/pkg/strings/split/key_value"
	"github.com/funtimecoding/go-library/pkg/system"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"strings"
)

func (c *Client) loadService() *gmail.Service {
	token := c.tokenFlow(c.credential, false)
	result, e := gmail.NewService(
		c.context,
		option.WithTokenSource(
			c.credential.TokenSource(c.context, token),
		),
	)
	errors.PanicOnError(e)
	name, _ := key_value.At(c.profile(result).EmailAddress)
	c.saveToken(
		token,
		system.Join(
			c.directory,
			fmt.Sprintf(
				"%s%s",
				strings.ToLower(name),
				constant.TokenSuffix,
			),
		),
	)

	return result
}
