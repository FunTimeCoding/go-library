package gmail

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gmail/constant"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/join"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

func (c *Client) loadCredential() *oauth2.Config {
	result, e := google.ConfigFromJSON(
		system.ReadBytes(join.Absolute(c.directory, constant.CredentialFile)),
		gmail.GmailReadonlyScope,
	)
	errors.PanicOnError(e)

	return result
}
