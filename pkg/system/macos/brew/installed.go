package brew

import (
	"github.com/funtimecoding/go-library/pkg/system/macos/brew/constant"
	"github.com/funtimecoding/go-library/pkg/system/macos/brew/response"
	"github.com/funtimecoding/go-library/pkg/system/run"
)

func (c *Client) Installed() *response.Installed {
	r := run.New()
	r.Start(
		constant.Brew,
		constant.Info,
		constant.Installed,
		constant.Notation2,
	)
	r.Verbose = true
	var result response.Installed
	r.ParseNotation(&result)

	return &result
}
