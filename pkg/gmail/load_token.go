package gmail

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"golang.org/x/oauth2"
)

func (c *Client) loadToken(path string) *oauth2.Token {
	f := system.Open(path)
	defer errors.PanicClose(f)
	result := &oauth2.Token{}
	errors.PanicOnError(json.NewDecoder(f).Decode(result))

	return result
}
