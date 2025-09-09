package brew

import (
	"github.com/funtimecoding/go-library/pkg/macos/brew/constant"
	"github.com/funtimecoding/go-library/pkg/macos/brew/response"
	"github.com/funtimecoding/go-library/pkg/system/run"
)

func (c *Client) Outdated() response.Outdated {
	r := run.New()
	r.Start(constant.Brew, constant.Outdated, constant.Notation)
	var result response.Outdated
	r.ParseNotation(&result)

	return result
}
