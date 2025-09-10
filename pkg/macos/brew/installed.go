package brew

import (
	"github.com/funtimecoding/go-library/pkg/macos/brew/constant"
	"github.com/funtimecoding/go-library/pkg/macos/brew/response"
	"github.com/funtimecoding/go-library/pkg/system/run"
)

func (c *Client) Installed() *response.Installed {
	r := run.New()
	r.Start(
		constant.Brew,
		constant.Info,
		constant.Installed,
		constant.Notation1,
	)
	r.Verbose = true
	var result response.Installed
	r.ParseNotation(&result)

	return &result
}
