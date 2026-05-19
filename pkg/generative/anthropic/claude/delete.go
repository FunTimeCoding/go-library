package claude

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"os"
	"path/filepath"
)

func (c *Client) Delete(sessionIdentifier string) {
	errors.PanicOnError(
		os.Remove(
			filepath.Join(
				c.base,
				join.Empty(sessionIdentifier, constant.NotationLogExtension),
			),
		),
	)
}
